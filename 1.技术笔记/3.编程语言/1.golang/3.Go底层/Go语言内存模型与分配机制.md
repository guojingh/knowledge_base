# Go语言内存模型与分配机制

**基于Go源码版本 1.19**

https://www.bilibili.com/video/BV1bv411c7bp?spm_id_from=333.788.player.switch&vd_source=15bae680e23049d4417233561746ed84&p=3

## 1.内存模型

### 1.1 操作系统存储模型

![image-20241114092304997](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114092304997.png)

观察上图，我们可以从中捕捉到关键词是：

- 多级模型
- 动态切换

### 1.2 虚拟内存与物理内存

![image-20241114092512619](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114092512619.png)

操作系统内存管理中，另一个重要概念是虚拟内存，其作用如下：

- 在用户与硬盘之间添加中间代理层（没有什么是加一个中间层解决不了的）
- 优化用户体验（进程感知到获得的内存空间时连续的）
- “放大”可用内存（虚拟内存可以由物理内存+磁盘补足，并根据冷热动态置换，用户无感知）

### 1.3 分页管理

操行系统中通常会将虚拟内存和物理内存切割成固定的尺寸，于虚拟内存而言叫做“页”，于物理内存而言叫做“帧”，原因及要点如下：

- 提高内存空间利用（以页为颗粒度后，消灭了不稳定的外部碎片，取而代之的是相对可控的内部碎片）
- 提高内外存交换概率（更大的颗粒度带来了更高的灵活度）
- 与虚拟内存机制呼应，便于建立虚拟内存 -> 物理内存的映射关系（聚合映射关系的数据结构，称为页表）
- linux 页/帧的大小固定，为4KB（这实际上是由实践推动的经验值，太粗会增加碎片率，太细会增加分配频率影响效率）

### 1.4 Golang 内存模型

Golang的内存模型设计的几个核心要点：

- 以空间换时间，一次缓存，多次复用

  由于每次向操作系统申请内存的操作很重，不妨一次多申请一些，以备后用。

  Golang 中的堆 mheap 正是基于该思想，产生的数据结构，我们可以从两个角度来解决 Golang 运行时的堆：

  - 对于操作系统而言，这是用户进程中缓存的内存。
  - 对于 Go 进程内部，堆是所有对象的内存起源。

- 多级缓存，实现无/细锁化

  ![image-20241114094445218](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114094445218.png)

堆是 Go 运行时中最大的临界共享资源，这意味着每次存取都需要加锁，在性能层面是一件很可怕的事。

为了解决这个问题，Golang 在堆 mheap 上，依次细化粒度，建立了 mcentral，mcache的模型，下面对三者进行一个梳理：

​	1.mheap：全局的内存起源，访问要加全局锁

​	2.mcentral：每种对象大小规格（全局共划分为 68 种）对应的缓存，锁的粒度也仅限于同一种规格以内。

​	3.mcache：每个 P （正是 GMP 中的 P）持有一种的内存模型，访问时无锁。

- 多级规格，提高效率

​	![image-20241114095457188](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114095457188.png)

首先理下 page 和 mspan 两个概念：

​	1.page：最小存储单元

​		Golang 借鉴操作系统分页管理的思想，每个最小的存储单元也称之为页 page，但大小为 8KB

​	2.mspan：最小的管理单元

​		mspan 大小为 page 的整数倍，且从 8 KB 到 80KB 被划分为 67 中不同的规格，分配对象时会根据大小映射到不同规格的 mspan，从中获取空间。

于是，mspan 产生了以下特点：

1. 根据规格大小，产生了等级的制度
2. 消除了外部碎片，但不可避免会有外部碎片
3. 宏观上能提高整体空间的利用率
4. 正是有了规格等级的概念，才支持 mcentral 实现细锁化

- 全局总览

  ![image-20241114100521993](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114100521993.png)

  ## 2.核心概念梳理

  ### 2.1 内存单元 mspan

  ![image-20241114100704230](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114100704230.png)

分点阐述 mspan 的特质：

- mspan 是 Golang 内存管理的最小单元

- mspan 大小是 page 的整数倍（Go 中的 page 大小为 8KB），且内部的页是连续的（至少在虚拟内存的视角中是这样）

- 每个 mspan 根据空间大小以及面向分配对象的大小，会被划分为不同的等级

- 由于同等级的 mspan 内聚于同一个 mcentral，最终会被组织成链表，因此带有前后指针（prev,next），所以会基于同一把互斥锁管理。

- mspan 会基于 bitMap 辅助快速找到空闲内存块（块大小为对应等级下的 object 大小），此时需要使用到 `Ctz64` 算法。

  ![image-20241114101807757](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114101807757.png)

mspan 类的源码位于 runtime/mheap.go 文件中：

```go
type mspan struct {
    _    sys.NotInHeap
    // 标识前后节点的指针
    next *mspan    
    prev *mspan     
    list *mSpanList

    // 起始地址
    startAddr uintptr
    // 包含几页，页是连续的
    npages    uintptr 
    manualFreeList gclinkptr
    // 标识此前的位置都已被占用
    freeindex uint16
   // 最多可以存放多少个 object
    nelems uint16 
    freeIndexForScan uint16
    // bitmap 每个 bit 对应一个 object 块，标识该块是否以被占用
    allocCache uint64

    // 标识 mspan 等级，包含 class 和 noscan 两部分信息
    spanclass             spanClass     // size class and noscan (uint8)
}
```

### 2.2 内存单元等级 spanClass

mspan 根据空间大小和面向对象分配的大小，被划分为 67 种等级（1-67，实际上还有一种隐藏的 0 级，用于处理更大的对象，上不封顶）

下表展示部分的 mspan 等级列表，数据取自 runtime/sizeclasses.go

![image-20241114103311486](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114103311486.png)

对上表各列进行解释：

1. class：mspan 等级标识，1-67

2. bytes/obj：该大小规格的对象会从这一 mspan 中获取空间，创建对象过程中，大小会向上取整为 8B 的整数倍，因此该表可以直接实现 object 到 mspan 的等级的映射。

3. bytes / span：该等级的 mspan 的总空间大小

4. object：该等级的 mspan最多可以 new 多少个对象，结果等于 (3)/(2)

5. tail waste：(3)/(2) 可能除不尽，于是该值为 (3)%(2)

6. max waste：通过下满来解释

   1. ![image-20241114104312261](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114104312261.png)

   2. 以 class 的 mpan 为例，class 分配的 object 大小统一为 24B，由于 object 大小 <= 16B 的会被分配到 class 2 以及之前的 class 中，因此自有 17B-24B 大小的 object 会被分配到 class 3。

      最不利的情况就是，当 object 大小为 17B，会残生浪费空间比例如下：

      ```sh
      ((24 - 17) * 341 + 8) / 8192 = 0.291358 = 29.24%
      ```

除了上面谈及的根据大小确定的 mspan 等级外，每个 object 还有一个重要的属性叫做noscan，标识了 object 是否包含指针，在 gc 时是否需要展开标记.
在Golang 中，会将 span class + noscan 两部分信息组装成一个 uint8，形成完整的
spanClass 标识.8 个 bit 中，高 7 位表示了上表的 span 等级 (总共 67 +1个等级，8个
bit 足够用了)，最低位表示 noscan 信息.

```go
type spanClass uint8

const (
    numSpanClasses = _NumSizeClasses << 1
    tinySpanClass  = spanClass(tinySizeClass<<1 | 1)
)

// uint8 左 7 位为 mspan 等级，最右一位标识是否为 noscan
func makeSpanClass(sizeclass uint8, noscan bool) spanClass {
    return spanClass(sizeclass<<1) | spanClass(bool2int(noscan))
}

//go:nosplit
func (sc spanClass) sizeclass() int8 {
    return int8(sc >> 1)
}

//go:nosplit
func (sc spanClass) noscan() bool {
    return sc&1 != 0
}
```

### 2.3 线程缓存 mcache

![image-20241114111716053](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114111716053.png)

1. mcache 是每个P都有的缓存，因此交互无锁
2. mcache 将每种 spanClass 等级的 mspan 各缓存了一个，总数为2（noscan维度） * 68 （维度大小）= 136
3. mcache 中还有一个为对象分配器 tiny allocator，用于处理小于16B对象的内存分配。在 3.3小节中详细展开。

```go
type mcache struct {
    _ sys.NotInHeap
    nextSample uintptr // trigger heap sample after allocating this many bytes
    scanAlloc  uintptr // bytes of scannable heap allocated

    // 微对象分配相关
    tiny       uintptr
    tinyoffset uintptr
    tinyAllocs uintptr

   // machce 中缓存的 mspan，每种 spanClass 各一个
    alloc [numSpanClasses]*mspan // spans to allocate from, indexed by spanClass
    stackcache [_NumStackOrders]stackfreelist
    flushGen atomic.Uint32
}
```

### 2.4 中心缓存 mcentral

![image-20241114113157050](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114113157050.png)

要点：

1. 每个 mcentral 对应一种 spanclass

2. 每个 mcentral 下聚合该 spanClass 下的 mspan

3. mcentral 下的 mspan 分为两个链表，分别为有空间 mspan 链表 partial 和满空间 mspan 链表 full

4. 每个 mcentral 一把锁

   ```go
   // Central list of free objects of a given size.
   type mcentral struct {
       _         sys.NotInHeap
       // 对应的 spanClass
       spanclass spanClass
       // 有空位的 mspan 集合，数组长度为 2 是用于抗一轮 GC
       partial [2]spanSet // list of spans with a free object
       // 无空位的 mspan 集合
       full    [2]spanSet // list of spans with no free objects
   }
   ```

### 2.5 全局堆缓存 mheap

要点：

- 对于 Golang 上层应用而言，堆是操作系统虚拟内存的抽象
- 以页（8KB）为单位，作为最小的存储单元
- 负责将连续页组装成 mspan
- 全局内存基于 bitMap 标识其使用情况，每个 bit 对应一页，为 0 则自有，为 1 则已被mspan组装
- 通过 heapArena 聚合页，记录到也到 mspan 的映射信息（2.7小节展开）
- 建立空闲页基数树索引 radix tree index，辅助快速寻找空闲页（2.6节展开）
- 是 mcentral 的持有者，持有所有 spanClass 下的 mcentral ，作为自身的缓存。
- 内存不够时，向操作系统申请，申请单位为 heapArena（64M）

## 3 对象分配流程

下面串联 Golang 中分配对象的流程，不论是一下哪种方式，最终都会殊途同归步入 mallocgc 方法中，并根据 3.1 小节中的策略执行分配流程：

- new(T)
- &T{}
- make(xxxx)

### 3.1 分配流程总览

Golang中，以及 object 的大小，将其分为下诉三类：

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114160133778.png" alt="image-20241114160133778" style="zoom:50%;" />

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114160149727.png" alt="image-20241114160149727" style="zoom:50%;" />

不同类型的对象，会有着不同的分配策略，这些内容在 mallocgc 方法中都有体现。

核心流程类似于读多级缓存的过程，由上而下，每一步只要成功则直接返回，若失败则由下层方法兜底。

**对于微对象的分配流程：**

1. 从 P 专属的 mcache 的 tiny 分配器取内存（无锁）
2. 根据所属的spanClass，从 P 专享 mcache 缓存的 mspan 中取内存（无锁）
3. 根据所属的 spanClass ，从对应的 mcentral 中取 mspan 填充到 mcache，然后从 mspan 中取内存（spanClass 锁粒度） 
4. 根据所属的 spanClass，从 mheap 的页分配器 pageAlloc 取得足够数量空闲页组装成 mspan，填充到 mcache，然后从 mspan 中取内存（全局锁）
5. mheap 向操作系统申请内存，更新页分配器的索引信息，然后重复（4）

对于小对象的分配刘晨是跳过（1）步，执行上述流程的 （2）-> （6）步

对于大对象的分配流程是跳过 （1）-> （3）步，执行上述流程的 （4）-> （5）步

![image-20241114161644714](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114161644714.png)



### 3.2 主干方法 mallocgc

代码位于 runtime/malloc.go

### 3.3 步骤1：tiny 分配

每个 P 独有的 mache 会有个微对象分配器，基于 offset 线性移动的方式对微对象进行分配，每 16B 成块，对象依据其大小，会向上取整为 2 的整数次幂进行空间补齐，然后进入分配流程。

 <img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20241114163154824.png" alt="image-20241114163154824" style="zoom: 80%;" />

### 3.4 步骤2：mcache 分配

### 3.5 步骤3：mcentral 分配





# 垃圾回收原理&源码走读

https://mp.weixin.qq.com/s/TdekaMjlf_kk_ReyPvoXiQ

https://mp.weixin.qq.com/s/Db19tKNer8D6FX6UG-Yujw











