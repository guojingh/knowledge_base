# remote-ssh 插件

## 1.介绍

remote-ssh 插件是 vs code 里面的一款插件，其用途在于远程开发。能让我们在 windows 的环境下很好的开发 linux 程序，上手体验比较良好，个人觉得比 Interactive authentication required（WSL）略胜一筹，配置简单，操作方便。

## 2.使用

### 1.下载插件

在 vs code 插件商店里面下载 remote-ssh 

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410173108272.png" alt="image-20240410173108272" style="zoom:67%;" />

下载完成之后，左侧会显示这个东西 ![image-20240410173247853](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410173247853.png)

### 2.配置

1. 点击左下角远程配置窗口 ![image-20240410173342098](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410173342098.png)
2. 点击连接到主机![image-20240410173411043](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410173411043.png)
3. 点击添加新的 SSH 主机 ![image-20240410173942145](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410173942145.png)
4. 输入 SSH 命令 这里注意用户名别输错了，我这里是 root，按回车确定![image-20240410174037583](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410174037583.png)
5. 指定 SSH 配置文件（随便制定推荐 home 目录）![image-20240410173818364](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410173818364.png)
6. 再次点击左下角，单机连接到主机![image-20240410174217719](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410174217719.png)
7. 点击你要连接的远程主机![image-20240410174413095](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410174413095.png)
8. 第一次连接远程主机会初始化 VS Code 服务器![image-20240410174616335](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410174616335.png)
9. 经过输入是否确定连接会显示输入密码![image-20240410181415614](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410181415614.png)
10. 当显示这个界面的时候，说明连上了，可以看到服务器上面的一些文件![image-20240410181540352](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410181540352.png)

### 3.后续配置

#### 1.配置秘钥（实现免密登陆）

```bash
# 生成公钥命令
ssh-keygen -t rsa -C "xxx@xxx.com"
```

在本地主机生成秘钥

![image-20240410183551917](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410183551917.png)

#### 2.将公钥上传到远程主机上

~~~bash
# 将公钥上传到指定的服务器 home 目录下，这里的路径要写全
PS C:\Users\guojinghu> scp C:\Users\guojinghu\.ssh\id_rsa.pub root@192.168.222.135:~
# 输出密码
root@192.168.222.135's password: 
id_rsa.pub    
~~~

#### 3.远程主机配置秘钥

远程主机生成一对秘钥，这里的邮箱和 windows 要保持一样

~~~bash
 ssh-keygen -t rsa -C "xxx@xxx.com"
~~~

![image-20240410184239452](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240410184239452.png)

#### 4.将 win 生成的秘钥追加到一个文件中

~~~bash
cat id_rsa.pub >> ~/.ssh/authorized_keys

root@master:~# cat id_rsa.pub >> ~/.ssh/authorized_keys
root@master:~# cd .ssh/
root@master:~/.ssh# ls
authorized_keys  id_rsa  id_rsa.pub
root@master:~/.ssh# cat authorized_keys 
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQD32geHdaG9DmnVo3Oxx5wJVCSZZwx1zil0nAOdOSIo1CGIdgfMkt8yb/CbU3/W2tQuBVNZ+78IZd8DFQh0gAO8gHDI61PiaKEO3cbAauN4PiTT2GGfroSYjAJXOytLcdHYBh1xzLJLY08wTSqBBAmTLJEXx9eR4j9D2WisW1IDvIJ9xhhDAN161uvxmODRSHKUiABzU7yzXgnz8JOA4SHawPhMlIJ4T0rRBdydPm+V8Alo7SO+FZ6X31K9sa6ZqRN/Gsk8OVgK5s3+3D65KDP2ytcvfwOtSz1+q1Z9ZApSR5s/lPSr+Ci8ViVlt9n9CbXKge31sv3hsWRFgue2LU+fyKylN/crkObWRK1jXxbIKvesfiBWjFzREs1OaSvLl9zHHRmgVQ8n/Kz0jWaojzd+NEj7dtTQxXjbVrt9FW+vPbQUyyVawHxNT2qHiqYGE9583j42Gx++BgsWhA2Sx1oslDVfmA6YvzHISb2cTT7WFSrFj2r96SmfsS2j4mD7XBU= xxx@xxx.com
~~~

那么这样的话不需要密码也能直接登陆了。

#### 5.图形化ubuntu 要进入服务器模式

~~~bash
# 设置多用户模式---最长用的服务器模式
root@master:~/.ssh# sudo systemctl set-default multi-user.target 
Created symlink /etc/systemd/system/default.target → /lib/systemd/system/multi-user.target.

# 返回图形化模式
root@master:~/.ssh# sudo systemctl set-default graphical.target 
Removed /etc/systemd/system/default.target.
Created symlink /etc/systemd/system/default.target → /lib/systemd/system/graphical.target.
~~~

#### 6.设置主机启动就进入服务器模式

~~~bash
# 编辑该文件
ahu@ahu:~/Documents$ sudo vim /etc/systemd/system/getty.target.wants/getty@tty1.service

# 将这里修改一下 
# ExecStart=-/sbin/agetty -o '-p -- \\u' --noclear %I $TERM  #原来的
ExecStart=-/sbin/agetty -a ahu --noclear %I $TERM   #修改后的结果
~~~





这样的话 remote-ssh 就配置完毕了！！！