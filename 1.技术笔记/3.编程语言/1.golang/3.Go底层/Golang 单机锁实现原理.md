# Golang 单机锁实现原理

## 0.前言

- 第一部分谈及 golang 最常用的互斥锁 sync.Mutex 的实现原理
- 第二部分则是以 Mutex 为基础，进一步介绍读写锁 sync.RWMutex 的时间原理

## 1.Sync.Mute

### 1.1 Mutex 核心机制

#### 1.1.1 上锁/解锁

遵循由简入繁的思路，我们首先忽略大量的实现细节以及基于并发安全角度的逻辑考量，思考实现一把锁最简单纯粹的主干流程：

- 通过 Mutex 内一个状态值标识锁的状态，例如，取 0 表示未加锁，1 表示已加锁；
- 上锁：把 0 改为 1；
- 解锁：把 1 置为 0.
- 上锁时，假若已经是 1，则上锁失败，需要等他人解锁，将状态改为 0.

Mutex 整体流程的骨架便是如此，接下来，我们就不断填充血肉、丰富细节.

#### 1.1.2 由自旋到阻塞的升级过程

一个优先的工具需要具备探测并适应环境，从而采取不同对策因地制宜的能力.

针对 goroutine 加锁时发现锁已被抢占的这种情形，此时摆在面前的策略有如下两种：

- 阻塞/唤醒：将当前 goroutine 阻塞挂起，直到锁被释放后，以回调的方式将阻塞 goroutine 重新唤醒，进行锁争夺；
- 自旋 + CAS：基于自旋结合 CAS 的方式，重复校验锁的状态并尝试获取锁，始终把主动权握在手中.

上述方案各有优劣，且有其适用的场景：

![image-20241125110936503](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241125110936503.png)

sync.Mutex 结合两种方案的使用场景，制定了一个锁升级的过程，反映了面对并发环境通过持续试探逐渐由乐观逐渐转换为悲观的状态，具体方案如下：

- 首先保持乐观，goroutine 采用自旋 + CAS 的策略争夺锁；
- 尝试持续受挫达到一定条件后，判定当前过于激烈，则由自旋转化为 阻塞/挂起模式。

上面谈及到的由自旋模式转为阻塞模式的具体条件拆解如下：

- 自旋累计达到4次仍未取得战果；
- CPU 单核或仅有单个 P 调度器；（此时自旋，其他 goroutine 根本没有机会释放锁，自旋纯属空转）
- 当前 P 的执行队列中仍有待执行的 G （避免因自旋影响到 GMP 调度效率）

#### 1.1.3 饥饿模式

1.1.2 小节的升级策略主要面向性能问题. 本小节引入的【饥饿模式】概念，则是展开对【公平性】的问题探讨.

下面首先拎清两个概念：

- 饥饿：顾名思义，是因为非公平机制的原因，导致 Mutex 阻塞队列中存在 goroutine 长时间取不到锁，从而陷入饥荒状态；
- 饥饿模式：当 Mutex 阻塞队列中存在处于饥饿态的 goroutine 时，会进入模式，将抢锁流程由非公平机制转为公平机制.

在 sync.Mutex 运行过程中存在两种模式：

- 正常模式/非饥饿模式：这是 sync.Mutex 默认采用的模式，当有 goroutine 从阻塞队列被唤醒时，会和此时先进入抢锁流程的 groutine 进行锁资源的争夺，假如抢锁失败，会重新回到阻塞队列头部。

  （值得一提的是，此时被唤醒的老 goroutine 相比新 goroutine 是处于劣势地位，因为新 goroutine 已经在占用 CPU 时间片，且新 goroutine 可能存在多个，从而形成多对一的人数优势，因此形势对老 goroutine 不利.）

- 饥饿模式：这是 sync.Mutex 为拯救陷入饥荒的老 goroutine 而启用的特殊机制，饥饿模式下，锁的所有权按照阻塞队列的顺序进行依次传递，新 goroutine 进行流程时不得抢锁，而是进入队列尾部排队。

两种模式的转换条件：

- 默认为正常模式
- 正常模式 -> 饥饿模式：当阻塞队列存在 goroutine 等锁超过 1ms 而不得，则进入饥饿模式；
- 饥饿模式 -> 正常模式：当阻塞队列已清空，或取得锁的 goroutine 等锁时间已低于 1ms时，则回到正常模式。

小结：正常模式灵活机动，性能好；饥饿模式严格死板，但等捍卫公平的底线。因此，两种模式的切换体现了 sync.Mutex 为适应环境变化，在公平与性能之间做出的调整与权衡，回头观望，这一项因地制宜，随机应变的能力是许多优秀工具所共有的特质。

#### 1.1.4 goroutine 唤醒标识

为尽可能缓解竞争压力和性能损耗，sync.Mutex 会不遗余力在可控范围内减少一些无意义的并发竞争和操作损耗.

在实现上，sync.Mutex 通过一个 mutexWoken 标识位，标志出当前是否已有 goroutine 在自旋抢锁或存在 goroutine 从阻塞队列中被唤醒；倘若 mutexWoken 为 true，且此时有解锁动作发生时，就没必要再额外唤醒阻塞的 goroutine 从而引起竞争内耗.

### 1.2 数据结构

```go
type Mutex struct {
    state int32
    sema  uint32
}
```

- state：锁中最核心的状态字段，不同 bit 位分别存储了 mutexLocked(是否上锁)、mutexWoken(是否有 goroutine 从阻塞队列中被唤醒)、mutexStarving(是否处于饥饿模式)的信息，具体在 1.2 节详细展开；
- sema：用于阻塞和唤醒 goroutine 的信号量。

#### 1.2.1 几个全局常量

```go
const (
    mutexLocked = 1 << iota // mutex is locked
    mutexWoken
    mutexStarving
    mutexWaiterShift = iota

    starvationThresholdNs = 1e6
)
```

• mutexLocked = 1：state 最右侧的一个 bit 位标志是否上锁，0-未上锁，1-已上锁；

• mutexWoken = 2：state 右数第二个 bit 位标志是否有 goroutine 从阻塞中被唤醒，0-没有，1-有；

• mutexStarving = 4：state 右数第三个 bit 位标志 Mutex 是否处于饥饿模式，0-非饥饿，1-饥饿；

• mutexWaiterShift = 3：右侧存在 3 个 bit 位标识特殊信息，分别为上述的 mutexLocked、mutexWoken、mutexStarving；

• starvationThresholdNs = 1 ms：sync.Mutex 进入饥饿模式的等待时间阈值.

#### 1.2.2 state 字段详述

Mutex.state 字段为 int32 类型，不同 bit 位具有不同标识含义：

![image-20241126105722927](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241126105722927.png)

低 3 位分别标识 mutexLocked（是否上锁）、mutexWoken（是否有协程在抢锁）、mutexStarving（是否处于饥饿模式），高 29 位的值聚合为一个范围为 0~2^29-1 的整数，表示在阻塞队列中等待的协程个数.

后续在加锁/解锁处理流程中，会频繁借助位运算从 Mutex.state 字段中快速获取到以上信息，大家可以先对以下几个式子混个眼熟：

- state & mutexLocked：判断是否上锁；
- state | mutexLocked：加锁；
- state & mutexWoken：判断是否存在抢锁的协程；
- state | mutexWoken：更新状态，标识存在抢锁的协程；
- state &^ mutexWoken：更新状态，标识不存在抢锁的协程；

(&^ 是一种较少见的位操作符，以 x &^ y 为例，假如 y = 1，结果为 0；假若 y = 0，结果为 x)

- state & mutexStarving：判断是否处于饥饿模式；
- state | mutexStarving：置为饥饿模式；
- state >> mutexWaiterShif：获取阻塞等待的协程数；
- state += 1 << mutexWaiterShif：阻塞等待的协程数 + 1.



















































