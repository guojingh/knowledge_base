"""
形象包平台监控 Demo —— 一个最小 FastAPI 服务，专门用来现场演示 PromQL。

暴露两类指标：
1. HTTP 指标（自动埋点）：http_requests_total / http_request_duration_seconds ...
   —— 由 prometheus-fastapi-instrumentator 自动生成，用来演示 QPS、错误率、P95 延迟。
2. 业务指标（手动埋点）：贴着分享文档里的形象包例子。
   - avatar_package_task_total   Counter   任务终态计数（成功率的分子分母）
   - package_stage_duration_seconds Histogram 各阶段耗时分布（P95）
   - avatar_queue_size           Gauge     当前排队任务数

配套接口用来"造数据"，方便边讲边点：
   GET /hello           快接口
   GET /slow            随机延迟，喂 Histogram
   GET /flaky           随机 500，喂错误率
   GET /task/train      模拟一次训练任务（随机成功/失败/超时）
   POST /queue/{n}      直接把队列长度设为 n，演示 Gauge
"""
import asyncio
import random
import time

from fastapi import FastAPI, HTTPException, Response
from prometheus_client import Counter, Gauge, Histogram
from prometheus_fastapi_instrumentator import Instrumentator

app = FastAPI(title="avatar-pkg-demo")

# ---------------------------------------------------------------------------
# 业务指标（手动埋点）—— 指标名与分享文档保持一致
# ---------------------------------------------------------------------------
package_task_total = Counter(
    "avatar_package_task_total",
    "形象包任务终态计数",
    labelnames=["path", "stage", "status"],
    # path:   unet | vis
    # stage:  training | conversion | vis_processing | register_figure
    # status: success | failed | timeout
)

package_stage_duration_seconds = Histogram(
    "package_stage_duration_seconds",
    "形象包各阶段耗时（秒）",
    labelnames=["path", "stage"],
    # 桶边界按业务目标设置：秒级任务，覆盖 0.5s ~ 30s
    buckets=(0.5, 1, 2, 5, 10, 20, 30),
)

queue_size = Gauge(
    "avatar_queue_size",
    "当前排队任务数",
)

# ---------------------------------------------------------------------------
# 接口
# ---------------------------------------------------------------------------
@app.get("/health")
def health():
    """健康检查接口：返回服务存活状态，不计入 HTTP 指标统计。"""
    return {"status": "ok"}


@app.get("/hello")
def hello():
    """快接口：稳定、低延迟，作为对照。"""
    return {"msg": "hello"}


@app.get("/slow")
async def slow():
    """随机延迟 10ms~1.5s，用来演示 http_request_duration_seconds 的分布与 P95。"""
    await asyncio.sleep(random.uniform(0.01, 1.5))
    return {"msg": "done"}


@app.get("/flaky")
def flaky():
    """约 20% 概率返回 500，用来演示 5xx 错误率查询。"""
    if random.random() < 0.2:
        raise HTTPException(status_code=500, detail="random failure")
    return {"msg": "ok"}


@app.get("/task/train")
def train(path: str = "unet"):
    """
    模拟一次训练任务：随机耗时 + 随机终态（成功 80% / 失败 15% / 超时 5%）。
    同时喂 Counter 和 Histogram，便于演示成功率与阶段耗时。
    """
    stage = "training"
    started = time.time()
    # 模拟任务耗时（这里用很短的 sleep 代替真实几十分钟）
    time.sleep(random.uniform(0.1, 0.8))
    duration = time.time() - started
    package_stage_duration_seconds.labels(path=path, stage=stage).observe(duration)

    roll = random.random()
    if roll < 0.80:
        status = "success"
    elif roll < 0.95:
        status = "failed"
    else:
        status = "timeout"
    package_task_total.labels(path=path, stage=stage, status=status).inc()
    return {"path": path, "stage": stage, "status": status, "duration": round(duration, 3)}


@app.post("/queue/{n}")
def set_queue(n: int):
    """直接设置队列长度，演示 Gauge 可增可减。"""
    queue_size.set(n)
    return {"queue_size": n}


# ---------------------------------------------------------------------------
# 暴露 /metrics（/health 与 /metrics 自身不计入 HTTP 统计）
# ---------------------------------------------------------------------------
Instrumentator(excluded_handlers=["/health", "/metrics"]).instrument(app).expose(app)
