// 零依赖 WebSocket echo 服务端，配合 index.html 使用
// 启动：node server.js，然后浏览器打开 http://localhost:8080
const http = require("http");
const fs = require("fs");
const path = require("path");
const crypto = require("crypto");

const PORT = 8080;
// RFC 6455 规定的固定 GUID，用于计算 Sec-WebSocket-Accept
const GUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11";

// 普通 HTTP 请求：返回 index.html
const server = http.createServer((req, res) => {
  fs.readFile(path.join(__dirname, "index.html"), (err, data) => {
    if (err) {
      res.writeHead(500);
      res.end("读取 index.html 失败");
      return;
    }
    res.writeHead(200, { "Content-Type": "text/html; charset=utf-8" });
    res.end(data);
  });
});

// 协议升级请求：完成 WebSocket 握手
server.on("upgrade", (req, socket) => {
  const key = req.headers["sec-websocket-key"];
  if (req.headers.upgrade?.toLowerCase() !== "websocket" || !key) {
    socket.end("HTTP/1.1 400 Bad Request\r\n\r\n");
    return;
  }
  const accept = crypto.createHash("sha1").update(key + GUID).digest("base64");
  socket.write(
    "HTTP/1.1 101 Switching Protocols\r\n" +
      "Upgrade: websocket\r\n" +
      "Connection: Upgrade\r\n" +
      `Sec-WebSocket-Accept: ${accept}\r\n\r\n`
  );
  console.log("客户端已连接");

  let buffer = Buffer.alloc(0);
  socket.on("data", (chunk) => {
    buffer = Buffer.concat([buffer, chunk]);
    // 一次 data 可能包含多个帧，也可能不足一帧
    let frame;
    while ((frame = parseFrame(buffer))) {
      buffer = buffer.subarray(frame.length);
      handleFrame(socket, frame);
    }
  });
  socket.on("close", () => console.log("客户端已断开"));
  socket.on("error", (err) => console.log("连接出错：", err.message));
});

// 解析一个数据帧，数据不足一帧时返回 null
function parseFrame(buf) {
  if (buf.length < 2) return null;
  const opcode = buf[0] & 0x0f;
  const masked = (buf[1] & 0x80) !== 0;
  let payloadLen = buf[1] & 0x7f;
  let offset = 2;

  if (payloadLen === 126) {
    if (buf.length < 4) return null;
    payloadLen = buf.readUInt16BE(2);
    offset = 4;
  } else if (payloadLen === 127) {
    if (buf.length < 10) return null;
    payloadLen = Number(buf.readBigUInt64BE(2));
    offset = 10;
  }

  // 客户端发给服务端的帧必须带 4 字节掩码
  const maskLen = masked ? 4 : 0;
  if (buf.length < offset + maskLen + payloadLen) return null;

  const mask = buf.subarray(offset, offset + maskLen);
  offset += maskLen;
  const payload = Buffer.from(buf.subarray(offset, offset + payloadLen));
  if (masked) {
    for (let i = 0; i < payload.length; i++) payload[i] ^= mask[i % 4];
  }
  return { opcode, payload, length: offset + payloadLen };
}

// 构造服务端发出的帧（服务端发送不加掩码）
function buildFrame(opcode, payload) {
  const len = payload.length;
  let header;
  if (len < 126) {
    header = Buffer.from([0x80 | opcode, len]);
  } else if (len < 65536) {
    header = Buffer.alloc(4);
    header[0] = 0x80 | opcode;
    header[1] = 126;
    header.writeUInt16BE(len, 2);
  } else {
    header = Buffer.alloc(10);
    header[0] = 0x80 | opcode;
    header[1] = 127;
    header.writeBigUInt64BE(BigInt(len), 2);
  }
  return Buffer.concat([header, payload]);
}

function handleFrame(socket, { opcode, payload }) {
  switch (opcode) {
    case 0x1: {
      // 文本帧：原样回传
      const text = payload.toString("utf8");
      console.log("收到消息：", text);
      socket.write(buildFrame(0x1, Buffer.from(text)));
      break;
    }
    case 0x8:
      // 关闭帧：回一个关闭帧后断开
      socket.end(buildFrame(0x8, payload));
      break;
    case 0x9:
      // ping：回 pong
      socket.write(buildFrame(0xa, payload));
      break;
    default:
      // 二进制帧、分片帧等本示例不处理
      break;
  }
}

server.listen(PORT, () => {
  console.log(`服务已启动：http://localhost:${PORT}`);
});
