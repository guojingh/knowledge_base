# Skill 实践指南

## Skill的结构与安装

### 目录结构

```python
my-skill/
	｜ - Skill.md    # 核心：description + 指令
	｜ - references/ # 详细参考文档
	｜ - scripts/    # 可执行脚本
```

### 安装方式

#### 方式一：Git仓库安装

```
npx skills add https://github.com/your-team/skills/main/code-review
```

从github仓库直接安装，支持社区共享的Skill。支持 CC 、Cursor 、WindSurf 等。

#### 方式二：手动放置安装

路径因平台而异，以 Claude Code 为例：

```shell
# 项目级
.claude/skills/code-review/SKILL.md
# 全局集
～/.claude/skills/code-review/SKILL.md
```

### SKILL.md 的结构

```markdown
---
name: my-skill
description: |
# 这里是 description 块
# 始终加载，用于出发判断
---
# 以下是body块（触发后才会加载）
## 工作流程
1.第一步...
2.第二步...
# 注意事项
- 关键约束
```

### Skill的工作原理

**STEP 01**

用户发起请求 --> 用户向agent发送任务指令

**STEP 02**

读取所有 description（agent读取已安装的所有skill的description，始终在上下文中）

**STEP 03**

语义匹配 --> 判断当前任务是否与某个Skill的描述相关

**STEP 04**

加载 SKILL.md body --> 命中后，将完整的Body内容注入到上下文

**STEP 05**

按SKILL执行指令，agent按照Skill提供的工作流和知识完成任务

### Description--激活的精准度

“Description是整个 skill 体系中最关键的一行字”

它决定了 Agent 在什么场景下会加载你的 Skill。写的不好就会导致两种致命的实效模式。

### Skill示例

SKILL.md

~~~markdown
---
name: get-china-weather
description: 使用wttr.in获取并显示指定地点的当前天气信息。当用户询问天气状况、天气预报或温度信息时，可以使用此功能
---

# 获取天气
使用wttr.in获取并显示指定地点的当前天气信息。

## 何时使用此SKILL

在以下情况使用此SKILL

- 用户询问当前天气情况
- 用户请求获取温度信息
- 用户需要天气预报
- 用户想知道特定城市或地点的天气情况
- 用户询问气象情况

## 流程

### 步骤1: 确定位置

从用户请求中提取位置信息

- 城市名称（例如：“伦敦”，“纽约”）
- 城市与国家（例如：“巴黎，法国”）
- 机场代码（例如：“LAX”）
- 坐标（例如：“~Eiffel+Towe”）

### 步骤2: 获取天气数据

使用curl从wttr.in获取天气信息：

```bash
curl "wttr.in/LOCATION?format=j1"
```

该服务返回的JSON格式数据如下：

- 当前状况
- 温度（摄氏度和华氏度）
- 湿度
- 风速和风向
- 天气描述
- 预测数据

### 步骤3: 解析并格式化输出

提取相关信息：

1.当前温度
2.天气状况（晴天，阴天，雨天等）
3.湿度百分比
4.风速
5."Feels like" temperature（体感温度）

### 步骤4: 展示结果

以清晰明了的格式展示天气信息

- 地点名称
- 当前状况
- 温度（摄氏度和华氏度）
- 其他细节（湿度，风）
- 如需，可提供简要预测

## 示例

### 示例1: 简单天气查询

**用户输入**
“西雅图的天气怎么样？”

**流程**

1.提取地点：“西雅图”
2.获取：`curl "wttr.in/Seattle?format=j1"`
3.解析JSON响应
4.格式化输出

**预期输出**
```text
西雅图天气：
当前：部分多云
温度：15°C（59°F）
感觉像：13°C (55°F)
湿度：65%
风速：10公里/小时，西北风
```
~~~
