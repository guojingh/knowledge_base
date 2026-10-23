# PromQL 演示环境（配合《技术分享-Prometheus监控实战》）

一个最小 FastAPI 服务 + Prometheus + Grafana + 自动造流量，用来现场演示 PromQL。
指标名和分享文档里的例子完全一致，文档 5.x 节标了 ▶ 的查询都能直接跑。

## 目录结构

```
demo/
├── app/
│   ├── main.py            # FastAPI 服务：自动 HTTP 指标 + 形象包业务指标
│   └── requirements.txt
├── Dockerfile
├── prometheus.yml         # 抓取配置，5s 一次
├── docker-compose.yml     # app + prometheus + grafana + loadgen
└── README.md
```

## 一键启动

```bash
cd demo
docker compose up -d --build
```

启动后：

- 应用指标：http://localhost:8100/metrics
- Prometheus UI（现场敲 PromQL 就用它）：http://localhost:9090
- Grafana（可选做看板）：http://localhost:3000
- `loadgen` 会自动持续打接口，30 秒后就有稳定曲线

停止 / 清理：

```bash
docker compose stop loadgen     # 停掉造流量，演示曲线掉下来 / 触发"无请求"
docker compose down             # 全部停掉并删除容器
```

## 接口说明（想手动造数据时用）

```bash
curl localhost:8100/hello              # 快接口，低延迟对照
curl localhost:8100/slow               # 随机 10ms~1.5s 延迟，喂 Histogram
curl localhost:8100/flaky              # 约 20% 概率 500，喂错误率
curl "localhost:8100/task/train?path=unet"   # 模拟训练任务：随机成功/失败/超时
curl -X POST localhost:8100/queue/42   # 把队列长度设为 42，演示 Gauge
```

## 演示脚本：按文档 5.x 顺序敲（复制到 Prometheus UI）

### 5.1 选指标 + 过滤

```promql
http_requests_total
http_requests_total{status=~"5.."}
```

### 5.3 Counter：rate / increase

```promql
# 每秒请求数（QPS）
sum(rate(http_requests_total[1m]))

# 最近 5 分钟总请求数
sum(increase(http_requests_total[5m]))
```

### 5.4 聚合：sum by

```promql
# 按状态码看 QPS
sum by (status) (rate(http_requests_total[1m]))

# 按接口看 QPS
sum by (handler) (rate(http_requests_total[1m]))
```

### 5.5 错误率（相除）

```promql
sum(rate(http_requests_total{status=~"5.."}[1m]))
/
sum(rate(http_requests_total[1m]))
```

> 打 `/flaky` 越多，这条曲线越明显。

### 5.6 Histogram：P95 延迟

```promql
# 全局 P95
histogram_quantile(0.95,
  sum by (le) (rate(http_request_duration_seconds_bucket[1m]))
)

# 分接口 P95（注意 le 必须保留）
histogram_quantile(0.95,
  sum by (handler, le) (rate(http_request_duration_seconds_bucket[1m]))
)
```

### 业务指标：形象包任务成功率 / 失败率

```promql
# 各终态每秒速率
sum by (status) (rate(avatar_package_task_total[1m]))

# unet 训练失败率（告警规则原型）
sum(rate(avatar_package_task_total{path="unet",stage="training",status="failed"}[5m]))
/
sum(rate(avatar_package_task_total{path="unet",stage="training"}[5m]))

# 各阶段 P95 耗时
histogram_quantile(0.95,
  sum by (path, stage, le) (rate(package_stage_duration_seconds_bucket[5m]))
)
```

### 5.7 状态判断

```promql
# 停掉 loadgen 后，这条会变成有结果（= 触发"无请求"）
sum(rate(http_requests_total[1m])) == 0

# 直接查 Gauge（先 POST /queue/120）
avatar_queue_size
avatar_queue_size > 100
```

## 不用 Docker 时的本地跑法

```bash
cd demo/app
pip install -r requirements.txt
uvicorn main:app --host 0.0.0.0 --port 8100
# 另开一个 Prometheus 指向 localhost:8100 即可
```
