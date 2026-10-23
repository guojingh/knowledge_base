"""FastAPI 入门示例

运行方式:
    fastapi dev main.py          # 开发模式（热重载）
    fastapi run main.py          # 生产模式

启动后访问:
    http://127.0.0.1:8000/           根路径
    http://127.0.0.1:8000/docs       自动生成的 Swagger 交互式文档
    http://127.0.0.1:8000/redoc      ReDoc 文档
"""
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

# 创建应用实例，应用的核心对象，所以路由都挂在它身上，
# title/version 不影响功能，只是 Swagger 文档的显示内容
app = FastAPI(title="FastAPI 示例", version="1.0.0")


# 内存中的简单数据存储，模拟数据库
items: dict[int, "Item"] = {}


class Item(BaseModel):
    """请求/响应的数据模型，Pydantic 会自动完成校验与序列化"""
    name: str
    price: float
    is_offer: bool | None = None


@app.get("/")
def read_root():
    """根接口，返回一条欢迎信息"""
    return {"message": "Hello FastAPI"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: str | None = None):
    """根据 id 查询单个 item，q 为可选查询参数"""
    if item_id not in items:
        raise HTTPException(status_code=404, detail="Item 不存在")
    return {"item_id": item_id, "item": items[item_id], "q": q}


@app.post("/items/{item_id}", status_code=201)
def create_item(item_id: int, item: Item):
    """创建一个 item，请求体会被自动校验为 Item 模型"""
    if item_id in items:
        raise HTTPException(status_code=400, detail="Item 已存在")
    items[item_id] = item
    return {"item_id": item_id, "item": item}
