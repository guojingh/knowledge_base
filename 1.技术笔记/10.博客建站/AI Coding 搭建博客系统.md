# AI Coding 搭建博客系统

## 



## 前端

### **技术选型**

Vue3.5

### **初始Prompt**

基于该vue版本设计一个博客系统的前端项目，代码文件放在 frontend 文件夹下：初始版本功能包括：在线编辑博客 管理博客（增，删，改，查） ，考虑到后需可能会添加其它功能，要注意项目的可扩展性，以及分层清晰性，注意注释标全一点。

****

### 项目架构设计

```
frontend/
├── index.html                     # 入口HTML
├── package.json                   # 依赖配置
├── vite.config.js                 # Vite构建配置
├── src/
│   ├── main.js                    # 应用入口
│   ├── App.vue                    # 根组件
│   ├── api/                       # 【API层】统一管理后端接口调用
│   │   └── blog.js                #   博客相关API
│   ├── stores/                    # 【状态管理层】Pinia全局状态
│   │   └── blog.js                #   博客状态Store
│   ├── router/                    # 【路由层】页面路由配置
│   │   └── index.js               #   路由定义
│   ├── views/                     # 【视图层】路由级别的页面组件
│   │   ├── Layout.vue             #   主布局框架
│   │   ├── BlogList.vue           #   博客列表页
│   │   ├── BlogDetail.vue         #   博客详情页
│   │   └── BlogEdit.vue           #   博客编辑/新建页
│   ├── components/                # 【组件层】可复用的UI组件
│   │   ├── BlogCard.vue           #   博客卡片
│   │   ├── BlogForm.vue           #   博客表单
│   │   └── ConfirmModal.vue       #   确认弹窗
│   └── utils/                     # 【工具层】通用工具函数
│       └── request.js             #   HTTP请求封装
```

**分层说明：**

- **Views** → 路由级页面，只负责组合组件和页面级逻辑
- **Components** → 可复用UI组件，通过 props/emits 通信
- **Stores (Pinia)** → 全局状态管理，管理跨组件共享数据
- **API** → 封装后端接口，Store 调用 API，View 调用 Store
- **Utils** → 纯工具函数，不依赖 Vue