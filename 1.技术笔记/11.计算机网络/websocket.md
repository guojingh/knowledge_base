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