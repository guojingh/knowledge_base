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

**短轮讯和WebSocket的区别：**

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

