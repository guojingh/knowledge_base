# WebSocket study

## 1.什么是WebSocket

### 1.1 WebSocket 诞生背景

早期，很多网站为了实现推送技术，所用的技术都是轮询（也叫短轮询）。轮询是指由浏览器每隔一段时间向服务器发出 HTTP 请求，然后服务器返回最新的数据给客户端。

**常见的轮询方式分为短轮询和长轮询，它们的区别如下：**

- **短轮询**

  短轮询的请求时间节点由客户端的定时器决定，和消息什么时候到达无关。

  ```mermaid
  sequenceDiagram
      participant C as 客户端
      participant S as 服务端
      C->>S: t=0  GET /poll
      S-->>C: 200 []（没有新数据）
      Note over C: 等定时器 10 秒
      C->>S: t=10 GET /poll
      S-->>C: 200 []（没有新数据）
      Note right of S: t=13 消息 A 到达，先存起来
      C->>S: t=20 GET /poll
      S-->>C: 200 [A]（晚了 7 秒）
      C->>S: t=30 GET /poll
      S-->>C: 200 []（没有新数据）
  ```



- **长轮询**

  长轮询客户端发下一个请求的时机，总是紧跟在上一个响应之后，所以同一时间只有一个请求在等待。

  ```mermaid
  sequenceDiagram
      participant C as 客户端
      participant S as 服务端
      C->>S: t=0 GET /lpoll
      activate S
      Note over S: 没有新数据，挂起等待
      Note right of S: t=13 消息 A 到达
      S-->>C: 200 [A]
      deactivate S
      C->>S: t=13 GET /lpoll
      activate S
      Note over S: 30 秒内无新消息
      S-->>C: 204 No Content（超时）
      deactivate S
      C->>S: t=43 GET /lpoll
      activate S
  ```

这种传统的模式带来很明显的缺点，即浏览器需要不断的向服务器发出请求，然而 HTTP 请求与响应可能会包含较长的头部，其中真正有效的数据可能只是很小的一部分，所以这样会消耗很多带宽资源。

> [!NOTE]
>
> 在这种情况下，HTML5定义了WebSocket协议，能更好地节省服务器资源和带宽，并且能够实时的进行通讯

WebSocket 使用 ws 或 wss 的统一资源标志符（URL），其中 wss 表示使用了 TLS 的 WebSocket。

如：ws://xxxx/websocket.org

**WebSocket与HTTP和HTTP使用相同的TCP端口，可以绕过大多数防火墙的限制**

**即：默认情况下：**

- WebSocket 协议使用 80 端口
- 若运行在 TLS 之上时，默认使用 443 端口

### 1.2 WebSocket 简介

WebSocket 是一种网络传输协议，可在单个 TCP 连接上进行全双工通信，位于 OSI 模型的应用层。WebSocket 协议在 2011 年由 IETF 标准化为 [RFC 6455](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Ftools.ietf.org%2Fhtml%2Frfc6455&objectId=1887095&objectType=1&contentType=undefined)，后由 [RFC 7936](https://cloud.tencent.com/developer/tools/blog-entry?target=https%3A%2F%2Ftools.ietf.org%2Fhtml%2Frfc7936&objectId=1887095&objectType=1&contentType=undefined) 补充规范。

WebSocket 使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据。在 WebSocket API 中，浏览器和服务器只需要完成一次握手，两者之间就可以创建持久性的连接，并进行双向数据传输。

**短轮询和WebSocket的区别：**

- 短轮询

  ```mermaid
  sequenceDiagram
      participant C as 客户端
      participant S as 服务端
  
      C->>S: t=0 GET /poll（带完整 HTTP 头）
      S-->>C: 200 []，没有新数据
      Note over C: 等定时器 10 秒
  
      C->>S: t=10 GET /poll
      S-->>C: 200 []，没有新数据
  
      Note right of S: t=13 消息 A 到达，只能先存起来
  
      C->>S: t=20 GET /poll
      S-->>C: 200 [A]，晚了 7 秒
  
      Note over C: t=25 要发消息 B
      C->>S: POST /send {B}（另发一个请求）
      S-->>C: 200 OK
  
      C->>S: t=30 GET /poll
      S-->>C: 200 []，没有新数据
  ```

- **WebSocket**

  ```mermaid
  sequenceDiagram
      participant C as 客户端
      participant S as 服务端
  
      C->>S: t=0 GET /ws，Upgrade: websocket
      S-->>C: 101 Switching Protocols
      Note over C,S: 握手完成，连接一直保持打开，之后只传 WebSocket 帧
  
      Note right of S: t=13 消息 A 到达
      S->>C: 文本帧 {A}，立即推送，没有延迟
  
      Note over C: t=25 要发消息 B
      C->>S: 文本帧 {B}，走同一条连接
  
      S->>C: ping
      C->>S: pong（心跳，保持连接）
  
      C->>S: close 帧
      S-->>C: close 帧
      Note over C,S: 连接关闭
  ```



### 1.3 WebSocket优点

1. 较少的控制开销：在连接创建后，服务器和客户端之间交换数据时，用于协议控制的数据包头部相对较小；
2. 更强的实时性：由于协议是全双工的，所以服务器可以随时主动给客户端下发数据。相对于 HTTP 请求需要等待客户端发起请求服务端才能响应，延迟明显更少；
3. 保持连接状态：与 HTTP 不同的是，WebSocket 需要先创建连接，这就使得其成为一种有状态的协议，之后通信时可以省略部分状态信息；
4. 更好的二进制支持：WebSocket 定义了二进制帧，相对 HTTP，可以更轻松地处理二进制内容；
5. 可以支持扩展：WebSocket 定义了扩展，用户可以扩展协议、实现部分自定义的子协议。

**由于 WebSocket 拥有上述的优点，所以它被广泛地应用在即时通讯/IM、实时音视频、在线教育、和游戏等领域。**

## 2.WebSocket API 

### 2.1 WebSocket的兼容性

![img](https://ask.qcloudimg.com/http-save/1037212/d277dccd786ed228b8839e3ff076da21.png)

由上图可知，目前主流的Web浏览器都支持WebSocket，所以目前基本上我们都所有项目里面都是可以使用WebSocket的。

使用 WebSocket  构造函数，我们就可以轻易地构造一个 WebSocket 对象。

**下面会从四个方面介绍 WebSocket API：**

1. WebSocket 的构造函数
2. WebSocket 对象的属性
3. WebSocket 的方法
4. WebSocket 的事件

### 2.2 构造函数

WebSocket 的构造函数为：

```ts
const myWebSocket = new WebSocket(url: string | URL, protocols?: string | string[]): WebSocket;
```

相关参数：

- URL：表示连接的URL，这是 WebSocket 服务器响应的 URL;
- protocols（可选）：一个协议字符串或者一个包含协议字符串的数组

**针对protocols**：这个字符串用于指定子协议，这样单个服务器可以实现多个 WebSocket 子协议。如果不指定协议字符串，则假定为空字符串。

### 2.3 属性

WebSocket 对象包含一下属性：

![image-20261119111626361](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20261119111626361.png)

**各属性具体含义：**

- binaryType：使用二进制的数据类型连接；
- bufferedAmount（只读）：未发送至服务器的字节数；
- extensions（只读）：服务器选择的扩展；
- onclose：用于指定连接关闭后的回调函数；
- onerror：用于指定连接失败后的回调函数；
- onmessage：用于指定当从服务器接受到信息时的回调函数；
- onopen：用于指定连接成功后的回调函数；
- protocol（只读）：用于返回服务器端选中的子协议的名字；
- readyState（只读）：返回当前 WebSocket 的连接状态，共有 4 种状态：
  - \- *CONNECTING — 正在连接中，对应的值为 0；*
  -   \- *OPEN — 已经连接并且可以通讯，对应的值为 1；*
  -   \- *CLOSING — 连接正在关闭，对应的值为 2；*
  -   \- *CLOSED — 连接已关闭或者没有连接成功，对应的值为 3*
- url（只读）：返回值为当构造函数创建 WebSocket 实例对象时 URL 的绝对路径。

### 2.4 方法

**WebSocket 主要有两个方法：**

```ts
// 该方法将需要通过 WebSocket 链接传输至服务器的数据排入队列，并根据所需要传输的数据的大小来增加 bufferedAmount 的值 。若数据无法传输（比如数据需要缓存而缓冲区已满）时，套接字会自行关闭。
send(data: string | ArrayBufferLike | Blob | ArrayBufferView): void;

// 该方法用于关闭 WebSocket 连接，如果连接已经关闭，则此方法不执行任何操作；
close(code?: number, reason?: string): void;
```

### 2.5 事件

使用 addEventListener()  来监听下面事件。

**以下是几个事件：**

- close：当一个 WebSocket 连接被关闭时触发，也可以通过 onclose 属性来设置；
- error：当一个 WebSocket 连接因错误而关闭时触发，也可以通过 onerror 属性来设置；
- message：当通过 WebSocket 收到数据时触发，也可以通过 onmessage 属性来设置；
- open：当一个 WebSocket 连接成功时触发，也可以通过 onopen 属性来设置

## 3. WebSocket 服务器

### 3.1 生命周期

![image-20261119154045094](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20261119154045094.png)

从上图可知，在使用 WebSocket 实现双全工通信之前，客户端与服务端之前需要进行握手（Handshake），在握手完成之后才能开始进行数据的双向通信。

**握手是在通信电路创建之后，信息传输开始之前。**

**握手用于达成参数，如：**

1. 信息传输率
2. 字母表
3. 奇偶校验
4. 中断过程
5. 其他协议特性

握手有助于不同结构的系统或设备在通信信道中连接，而不需要人为设置参数。

**分析WebSocket的握手协议**

### 3.2 握手协议

`WebSocket` 协议属于应用层协议，它依赖于传输层的 `TCP` 协议。WebSocket 通过 `HTTP/1.1` 协议的 `101 `状态码进行握手。为了创建 `WebSocket`连接，需要通过浏览器发出请求，之后服务器进行回应，这个过程通常称为"握手" （Handshaking）。

**注意：**

`HTTP/1.1` 是最经典的 `Upgrade` 握手，也是最常见的； `HTTP/2` `HTTP/3` 的握手方式有所不同。

| 协议     | 握手方式                     | RFC      |
| -------- | ---------------------------- | -------- |
| HTTP/1.1 | `Upgrade: websocket` + `101` | RFC 6455 |
| HTTP/2   | 扩展 `CONNECT` + `:protocol` | RFC 8441 |
| HTTP/3   | 扩展 `CONNECT` + `:protocol` | RFC 9220 |

**利用HTTP完成握手的几个好处：**

- 首先：让 `WebSocket` 与现有的 HTTP 基础设施兼容----使得 `WebSocket` 服务器可以运行在 `80` 和 `443` 端口上，这通常是对客户端唯一开放端口；
- 其次：让我们可以完全重用并扩展 `HTTP` 的 `Upgrade` 流，为其添加自定义的 `WebSocket`  首部，以完成协商。

#### 3.1.1 客户端请求

```json
GET ws://echo.websocket.org/ HTTP/1.1
Host: echo.websocket.org
Origin: file://
Connection: Upgrade
Upgrade: websocket
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: Zx8rNEkBE4xnwifpuh8DHQ==
Sec-WebSocket-Extensions: permessage-deflate; client_max_window_bits
```

**上述字段说明：**

- Connection：必须设置 Upgrade，表示客户端希望连接升级
- Upgrade：字段必须设置为 Websocket，表示希望升级为 WebSocket 协议
- Sec-WebSocket-Version：表示支持的 WebSocket 版本。RFC6455 要求使用的版本是 13，之前草案的版本均应当弃用；
- Sec-WebSocket-Key：是随机的字符串，服务器端会用这些数据来构造出一个 SHA-1 的信息摘要
- Sec-WebSocket-Extensions：用于协商本次连接要使用的 WebSocket 扩展：客户端发送支持的扩展，服务器通过返回相同的首部确认自己支持一个或多个扩展
- Origin：字段是可选的，通常用来表示在浏览器中发起此 WebSocket 连接所在的页面，类似于 Referer。但是，与 Referer 不同的是，Origin 只包含了协议和主机名称

**针对Sec-WebSocket-Key**：把 “*Sec-WebSocket-Key*” 加上一个特殊字符串 “*258EAFA5-E914-47DA-95CA-C5AB0DC85B11*”，然后计算 SHA-1 摘要，之后进行 Base64 编码，将结果做为 “*Sec-WebSocket-Accept*” 头的值，返回给客户端。如此操作，可以尽量避免普通 HTTP 请求被误认为 WebSocket 协议。

#### 3.1.2 服务端响应

```json
1.HTTP/1.1 101 Switching Protocols
2.Upgrade: websocket
3.Connection: Upgrade
4.Sec-WebSocket-Accept: vOCj83xdUaxls8CxBz34jP2+CW0=
```

**上述字段说明：**

1. 101 响应码确认升级到 WebSocket 协议
2. 设置 Connection 头的值为 “Upgrade” 来指示这是一个升级请求（HTTP 协议提供了一种特殊的机制，这一机制允许将一个已建立的连接升级成新的、不相容的协议）
3. Upgrade 头指定一项或多项协议名，按优先级排序，以逗号分隔。这里表示升级为 WebSocket 协议
4. 签名的键值验证协议支持

### 3.3 消息通信基础

在 `WebSocket` 协议中，数据是通过一系列数据帧来进行传输。

为了避免由于网络中介（例如一些拦截处理）或者一些安全问题，客户端必须在它发送到服务端的所有帧中添加掩码。服务端收到没有添加掩码的数据帧后，必须立即关闭连接。

#### 3.3.1 WebSocket 的数据帧格式

![img](https://ask.qcloudimg.com/http-save/1037212/9ce292bc8ad93137b6cdf227232db4bb.png)

**逐个字节看：**

- **第一个字节（最前面8位）**	

  | 字段         | 位数 | 含义                                                         |
  | ------------ | ---- | ------------------------------------------------------------ |
  | **FIN**      | 1    | 是不是消息的最后一帧。1=这是最后一帧；0=后面还有续帧（分片） |
  | **RSV1/2/3** | 各 1 | 保留位，一般是 0，除非用了扩展（比如压缩 `permessage-deflate`） |
  | **opcode**   | 4    | 帧类型。`0x1`=文本，`0x2`=二进制，`0x0`=续帧，`0x8`=关闭，`0x9`=ping，`0xA`=pong |

- **第二个字节**

  | 字段            | 位数 | 含义                                                         |
  | --------------- | ---- | ------------------------------------------------------------ |
  | **MASK**        | 1    | 负载是否加了掩码。**客户端发给服务端必须是 1**，服务端发给客户端必须是 0 |
  | **Payload len** | 7    | 负载长度（见下面的三段式）                                   |

- **长度为什么要分三段**

  - 值 `0~125`：就是真实长度，后面没有额外长度字段。
  - 值 `126`：真实长度放在**后面的 16 位**里（图中 "Extended payload length 16"），最大 65535。
  - 值 `127`：真实长度放在**后面的 64 位**里（图中那行 "continued, if payload len == 127"），用于超大消息。

- **Masking-key（掩码）**

  只有 MASK = 1 时才存在，占四个字节。客户端会用它做负载或异或处理，服务端收到后用同样的 `key` 处理或还原。

- **Payload Data（负载）**

  真正的消息内容，文本帧就是 UTF-8 字节。

#### 3.3.2 掩码算法

掩码字段是一个由客户端随机选择的 32 位的值。掩码值必须是不可被推测的。因此掩码必须来自强大的`熵源`,并且给定的掩码不能让服务器或者代理轻易推测到后续帧。掩码的不可预测性对于预防恶性应用的作者在网上暴露相关的字节数据至关重要。「掩码不影响数据荷载的长度，对数据进行掩码操作和对数据进行反掩码操作所涉及的步骤是相同的」。

在 `WebSocket` 协议中，数据掩码的作用是增强协议的安全性。但数据掩码并不是为了保护数据本身，因为算法本身是公开的，运算也不复杂。那么为什么还要引入数据掩码呢？引入数据掩码是为了防止早期版本的协议中存在的代理缓存污染攻击等问题。了解完 WebSocket 掩码算法和数据掩码的作用之后，我们再来介绍一下数据分片的概念。

#### 3.3.4 数据分片

`WebSocket` 的每条消息可能被切分成多个数据帧。当 `WebSocket` 的接收方收到一个数据帧时，会根据 FIN 值来判断，是否已经收到消息的最后一个数据帧。

利用 `FIN` 和 `Opcode` 我们可以跨帧发送消息。

**操作码告诉帧应该做什么？**

- 如果是 0x1，有效荷载就是文本；
- 如果是 0x2，有效荷载就是二进制数据；
- 如果是 0x0，则该帧是一个延续帧（这意味着服务器应该将帧的有效负载连接到从该客户机接收到的最后一个帧）

### 3.4 实现消息通信功能

参考 server.js 中的服务端实现



## 4. WebSocket 学习过程中的易错常识

### 4.1 WebSocket 与 HTTP 有什么关系

WebSocket 是一种与 HTTP 不同的协议。两者都位于 OSI 模型应用层，并且都依赖于传输层的 TCP 协议。

虽然他们不同，但是 RFC 6455 中规定：WebSocket 被设计为在 HTTP 80 和 443 端口上工作，并支持HTTP代理和中介，从而使其与HTTP协议兼容，为了实现兼容，WebSocket 握手使用 HTTP Upgrade 头，从 HTTP 协议更改为 WebSocket 协议。

既然已经提到了OSI（Open System Interconnection Model）模型，下面是一张很生动，形象描述 OSI 模型的示意图。

![image-20261124214906080](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20261124214906080.png)

### 4.2 什么是 WebSocket 心跳

网络中的接收和发送数据都是使用 Socket 进行实现。但是如果此套接字已经断开了，那发送数据和接收数据的时候就一定有问题。

**可是如何判断这个套接字是否可以使用呢？**这个就需要在系统中创建心跳机制。

所谓"心跳"就是定时发送一个自定义的结构体（心跳包或者心跳帧），让对方知道自己"在线"，以确保链接的有效性。

而所谓的心跳包就是客户端定时发送简单的信息给服务端告诉它我还在而已。代码就是每隔几分钟发送一个固定信息给服务端，服务端收到后回复一个固定信息，如果服务端几分钟内没有收到客户端信息则视为客户端断开。

**在 WebSocket 协议中定义了 心跳 Ping 和 心跳 Pong 的控制帧：**

- 心跳 Ping 帧包含的操作码是 0x9：如果收到了一个心跳 Ping 帧，那么终端必须发送一个心跳 Pong 帧作为回应，除非已经收到了一个关闭帧。否则终端应该尽快回复 Pong 帧；
- 心跳 Pong 帧包含的操作码是 0xA：作为回应发送的 Pong 帧必须完整携带 Ping 帧中传递过来的 “应用数据” 字段。

如果终端收到一个 Ping 帧但是没有发送 Pong 帧来回应之前的 Ping 帧，那么终端可以选择仅为最近处理的 Ping 帧发送 Pong 帧。此外，可以自动发送一个 Pong 帧，这用作单向心跳。

### 4.3 Socket 是什么？

网络上的两个程序通过一个双向的通信连接实现数据的交换，这个连接的一端称为一个 Socket（套接字），因此建立网络通信连接至少要一对端口号。

**Socket 本质：**是对 TCP/IP 协议栈的封装，它提供了一个针对 TCP 或者 UDP 编程的接口，并不是另一种协议。通过 Socket，你可以使用 TCP/IP 协议。

> [!IMPORTANT]
>
> **关于 Socket，可以总结以下几点：**
>
> - 它可以实现底层通信，几乎所有的应用层都是通过 socket 进行通信的；
> - 对 TCP/IP 协议进行封装，便于应用层协议调用，属于二者之间的中间抽象层；
> - TCP/IP 协议族中，传输层存在两种通用协议: TCP、UDP，两种协议不同，因为不同参数的 socket 实现过程也不一样。

**下图说明了面向连接的协议的套接字 API 的客户端/服务器关系：**

![img](https://ask.qcloudimg.com/http-save/1037212/4fa83e630b5210799050b53abcbcca3c.png)



