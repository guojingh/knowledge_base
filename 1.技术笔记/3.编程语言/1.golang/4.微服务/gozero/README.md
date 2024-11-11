## gozero 

### goctl 命令
```bash
## 生成全新项目
goctl api new shorturl

## 初始化 template
goctl template clean

## 根据指定 api 文件生成对应代码
goctl api go -api shorturl.api -dir . -style=goZero

```