# Charles抓包工具

## 1.安装Charles

官网地址

https://www.charlesproxy.com/download/

## 2.破解

### 2.1 生成密钥

Charles破解工具：https://www.zzzmode.com/mytools/charles/

**1.生成激活码**

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516111940975.png" alt="image-20250516111940975" style="zoom:50%;" />

点击生成

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516112214267.png" alt="image-20250516112214267" style="zoom:50%;" />

### 2.2 登录密钥

**1.打开Charles, Help -> Register Charlse**

![image-20250516112415055](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516112415055.png)

**2.输入 2.1 生成的name和密钥**

![image-20250516112505734](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516112505734.png)

**3.提交后显示成功**

![image-20250516112551275](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516112551275.png)

## 3.配置

### 3.1 证书安装-mac

选中Charles，在顶部工具栏中 **Help--SSL Porxying**

依次打开

![image-20250516112804341](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516112804341.png)

### 3.2 设置抓包类型，端口号

在Charles顶部工具栏 **Proxy--Proxy Setting**，可以设置抓包数据类型，包括http与socket数据。可以根据需要在proxies栏下勾选。

![image-20250516112855010](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516112855010.png)

### 3.3 抓取HTTPS请求的配置

HTTPS抓包，需要经过SSL。

Charles顶部工具栏 **Proxy--SSL Proxy Setting**

点击Add，Host栏与Port栏 **都为空**（表示抓所有SSL请求），确定即可。

![image-20250516113151348](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516113151348.png)

如果没有设置此项，面对HTTPS请求，会出现访问失败（请求左侧出现小锁头图标）

![image-20250516113215308](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516113215308.png)

## 4.连接手机

### 4.1 手机安装证书（安卓）

选中Charles，在顶部工具栏中 **Help--SSL Porxying**

点击 Charles Root Certificate on a Mobile Device or Remote Browser后出现弹框

![image-20250516113423409](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516113423409.png)

### 4.2 连接手机

手机配置WIFI代理，确保手机和PC端处在同一网络环境下

打开 手机WIFI -- 更多WIFI设置 -- 连接PC同一WIFI -- 配置代理

![883F67B9FA85837D9D984D953491A665](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/883F67B9FA85837D9D984D953491A665.png) 

如果是第一次连接的设备，会弹框询问是否允许，点击allow即可

![image-20250516114148095](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250516114148095.png)

### 4.3 手机下载并信任证书

手机打开默认浏览器，在地址拦输入 **chls.pro/ssl** 并访问，出现弹框询问是否安装，同意安装即可。

安装完成后，在 **设置 -- 安全 -- 更多安全设置 -- 从手机存储安装 -- CA证书 -- 任然安装 -- 选择刚下载的CA证书进行安装**



**然后就可以进行抓包了**



**注意：**

如果安装证书后，Charles 抓包仍然报以下错误：You may need to configure your browser or application to trust the Charles Root Certificate. See SSL Proxying in the Help menu.

**原因：**

这是 Android 7.0 的问题，在 7.0 上，用户自己安装的证书是不被 app 应用信任的(安全性考虑，防止 Charles 这种做接口逆向，监听数据)

**解决办法：**

1. 如果有自己的测试包那么安装测试包进行抓包
2. 使用安卓模拟器，可能会比比较卡顿
3. 添加root根CA证书（没进行尝试过）



