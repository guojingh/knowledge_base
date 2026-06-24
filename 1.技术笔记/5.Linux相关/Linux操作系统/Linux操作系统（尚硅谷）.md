# Linux操作系统（尚硅谷）

## 课程介绍

### **学了Linux可以做什么？**

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250310080633602.png" alt="image-20250310080633602" style="zoom:67%;" />

### 课程内容

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250310080814838.png" alt="image-20250310080814838" style="zoom:67%;"/>

### 课程特色

<img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250310081019902.png" alt="image-20250310081019902" style="zoom:67%;"/>

## Linux概述

- Linux是一个操作系统（OS）

  ![image-20250310082056746](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250310082056746.png)

### Linux的诞生

- 李纳斯·托瓦斯（Linux Torvalds）
  - 上大学期间，对Unix产生了浓厚兴趣
  - 1991年初，开始基于Minix(Unix的变种)进行开发
  - 1991年6月，确定开发一个类Unix操作系统的内核
  - 1991年9月，开发完成内核版本的 0.01 版本，命名为Linux

### Linux和Unix的渊源

![image-20250310084518857](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250310085454515.png)

### GNU/Linux

 <img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250310085454515.png" alt="image-20250310085454515" style="zoom:67%;" />

### Linux的发行版

 <img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250311074502969.png" alt="image-20250311074502969" style="zoom:67%;" />

### Linux vs Windows

![image-20250311074547471](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250311074547471.png)

## Linux 文件系统

### Linux 文件

Linux 系统中一切皆文件

### Linux 目录结构

![image-20250312081427532](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250312081427532.png)

- /bin
  - 是 Binary 的缩写，这个目录存放着最经常使用的命令
- /sbin
  - s 是 system 的意思，这里存放的是系统管理员使用的系统管理程序
-  /lib
  - 库目录，存放系统必要的第三方库
- /lib64
  - 库目录，存放系统必要的第三方库(64位)
- /usr
  - 存放用户级别的东西
- /boot
  - 存放系统启动时的引导程序
- /dev
  - 存放设备的相关信息（cpu，memory等）
- /etc
  - 配置文件
- /home
  - 用户的主目录
- /root
  - root用户的主目录
- /opt
  - 第三方软件存放目录
- /media
  - 识别外部储存设备（U盘，光驱的挂载点）
- /mnt
  - 功能和 /media 类似
- /proc  （process）
  - 系统进程的信息，对于系统比较重要
- /run
  - 存放系统运行的实时信息
- /srv
  - 存放系统服务相关的信息
- /sys
  - 存放系统硬件的相关信息
- /tmp
  - 临时目录，临时存在的东西
- /var
  - 可变目录，存放一些变量（各种各样的日志）

## VIM 编辑器

### vi/vim 是什么

VI 是 Unix 操作系统和类 Unix 操作系统中最通用的文本编辑器

VIM 编辑器是从 VI 发展出来的一个性能强大的文本编辑器。

### 一般模式

以 vim 打开一个档案就直接进入了一般模式了（这是默认的模式）

1. 常用语法

   ![image-20250313083413368](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250313083413368.png)

   ![](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250313084650341.png)

   

### 插入模式

- i：进入插入模式



### 指令模式

1. 基本语法

   ![image-20250314083231360](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314083231360.png)

### 模式间转换

![image-20250313082411137](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250313082411137.png)

### vi / vim 键盘图

![image-20250314084037144](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314084037144.png)

## 网络配置和系统管理操作

### 查看网络IP和网关

#### 1.查看虚拟网络编辑器

#### 2.VMware提供了三种网络连接模式

- 桥接模式

  虚拟机直接连接外部物理网络的模式，主机起到了网桥的作用，这种模式下，虚拟机可以直接访问外部网络，并且对外部网络是可见的。

- NAT模式

  虚拟机和主机构建一个专用网络，并通过虚拟网络地址转换（NAT）设备对 IP 进行转换，虚拟机通过共享主机IP可以访问外部网络，但外部网络无法访问虚拟机。

- 仅主机模式

  虚拟机只与主机共享一个专用网络，与外部网络无法通信。

#### 3.**修改虚拟机自动分配IP为静态IP**

1. 修改配置文件 /etc/sysconfig/net/ifcfg-ens33

   ![image-20250318224900756](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318224900756.png)

2. 重启网络服务

   ```sh
   service network restart
   ```

#### 4.配置主机名

1. 修改主机名

   第一种方式

   - 修改配置文件 /etc/hostname
   - 重启服务器

   第二种方式 hostnamectl 命令

   ```sh
   hostnamectl set-hostname spark10
   ```

   ![image-20250318230053773](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318230053773.png)

2. 添加主机名和IP的映射关系（hosts文件）

   ![image-20250318230620378](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318230620378.png)

3. 修改 windows 上的hosts文件，以便使用主机名进行访问

   ![image-20250318231107738](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318231107738.png)

   ![image-20250318231124701](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318231124701.png)

   这样就可以使用域名进行访问了

   ![image-20250318231305393](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318231305393.png)

#### 5.远程登陆

使用 XShell 工具

![image-20250320081454545](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250320081454545.png)

### 系统管理

#### 1.Linux 中的进程和服务

计算机中，一个正在执行的程序或命令，被叫做“进程”（process）.

启动之后，一直存在，常驻内存的进程，一般被称为“服务”（service）.

#### 2.service 服务管理 （Centos 6 版本-了解）

1. 基本语法

   service 服务名 start | stop | restart | status

2. 经验技巧

   查看服务的方法：/etc/init.d/服务名，发现只有两个服务保留在 service

   ![image-20250320081820064](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250320081820064.png)

#### 3. systemctl (Centos 7 版本-重点掌握)

1. 基本语法

   systemctl start | stop | restart | status 服务名

2. 经验技巧

   查看服务的方法： /usr/lib/systemd/system

   eg: 网络服务

   ![image-20250320082541865](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250320082541865.png)

   这两个都是 Centos 的网络服务，network 是老版，NetworkManager 是新版，那么在Linux系统中，保证其中一个在运行就 ok 了。

#### 4.systemctl 设置后台服务的自启配置

1. 控制台输入 setup 命令

2. 选择系统服务

   <img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250320084827862.png" alt="image-20250320084827862" style="zoom:50%;" />

3. 前面设置 * 的即是自启动服务，其中点击空格进行配置

   <img src="https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250320085010111.png" alt="image-20250320085010111" style="zoom: 80%;" />

#### 5.系统运行级别

1. Linux 运行级别（Centos6）

   ![image-20250320085117982](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250320085117982.png)

   查看默认级别： /vi/etc/inittab

   Linux系统有7种运行级别（runlevel）：常用的级别是3和5

   - 运行级别0：系统停机状态，系统默认运行级别不能设为0，否则不能正常启动
   - 运行级别1：单用户工作状态，root权限，用于系统维护，禁止远程登陆
   - 运行级别2：多用户状态（没有NFS）,不支持网络
   - 运行级别3：完全的多用户状态（有NFS），登陆后进入控制台命令行模式
   - 运行级别4：系统未使用，保留
   - 运行级别5：X11控制台，登陆后进入图形GUI模式
   - 运行级别6：系统正常关闭并重启，默认运行级别不能设为6，否则不能正常启动

2. Centos7 的运行级别简化为：

   multi-user.target 等价于原运行级别 3 （多用户有网，无图形化界面）

   graphical.target 等价于原运行级别 5 （多用户有网，有图形化界面）

3. 查看当前运行级别

   systemctl get-default

4. systemctl 命令

   `systemctl` 是 Linux 系统中用于管理系统和服务的命令行工具，主要用于使用 systemd 系统和服务管理器。

   1. 启动服务

      systemctl start <服务名>

   2. 停止服务

      systemctl stop <服务名>

   3. 重启服务

      systemctl restart <服务名>

   4. 查看服务状态

      systemctl status <服务名>

   5. 开机自启用服务

      systemctl enable <服务名>

   6. 取消开机自启用

      systemctl disable <服务名>

   7. 查看所有服务

      systemctl list-units --type=service

      systemctl list-units-files (查看所有服务开机自启动状态)

   8. 查看服务日志

      journalctl -u <服务名>

5. 关闭防火墙（Centos7）

   - systemctl status firewalld ---- 查看防火墙状态
   - systemctl stop firewalld ---- 停止防火墙
   - systemctl disable firewalld ---- 开机自动关闭防火墙

6. 关机重启命令

   在 Linux 领域内大多用在服务器上，很少遇到关机的操作。毕竟服务器上跑一个服务是永无止境的，除非特殊情况下，不得已才会关机。

   1. 基本语法

      1. sync --- 将数据有内存同步到硬盘中

      2. halt --- 停机，关闭系统，但不断电

      3. poweroff --- 关机，断电

      4. reboot --- 重启，等同于 shutdown -r now

      5. shutdown [选项] 时间

         ![image-20250323094236375](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250323094236375.png)


## 常用基本命令（重要）

Shell 可以看作是一个命令解释器，为我们提供了交互式的文本控制台界面。我们可以通过终端控制台来输入命令，由Shell 进行解释并最终交给内核执行。本章就将分类介绍常用的 shell 命令。

#### 1.帮助命令

##### 1.1 man 获取帮助信息

- 基本语法

  man[命令或配置文件]   ---   功能描述：获取帮助信息

- 显示说明

  ![image-20250323095758519](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250323095758519.png)

- 案例实操

  - 查看 ls 命令的帮助信息

    man ls

    ls --help --- 也可以获取外置命令使用信息

##### 1.2 help 获取 shell 内置命令的帮助信息

一部分基础功能的系统命令是直接内嵌在 shell 中的，系统加载启动之后会随着 shell 一起加载，常驻系统内存中。这部分命令被称为 “内置（bulit-in）命令”；相应的其他命令被称为“外部命令”。

- 基本用法

  help 命令   ---   功能描述：获取 shell 内置命令的帮助信息

- 案例实操

  查看 cd 命令的帮助信息

  help cd

##### 1.3 常用快捷键

![image-20250323101016891](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250323101016891.png)

#### 2.文件目录类

##### 2.1 pwd 显示当前工作目录的绝对路径

pwd : print working directory 打印工作目录

- 基本语法

  pwd   ---   显示当前工作目录的绝对路径

- 案例实操

  ![image-20250323101637534](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250323101637534.png)

##### 2.2 ls列出目录的内容

ls: list 列出目录内容

- 基本语法

  ls [选项] [目录或是文件]

- 选项说明

  ![image-20250324075905308](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324075905308.png)

- 显示说明

  每行列出的信息依次是：**文件类型和权限 链接数 文件属主 文件属组 文件大小用byte来表示 建立或最近修改的时间 名字**

- 实操案例

##### 2.3 创建文件夹

mkdir: 创建文件夹

mkdir -p a/b/c

##### 2.4 删除文件夹

rmdir: 删除文件夹 

rmdir -p a/b/c

##### 2.5 touch 创建空文件

- 基本语法

  touch 文件名称

- 案例实操

  touch xiyou/dssz/sunwukong.txt

##### 2.6 cp 复制文件或目录

- 基本语法

  cp [选项] source dest  ---   复制source文件到dest

- 选项说明

  ![image-20250324081549410](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324081549410.png)

- 参数说明

  ![image-20250324081908509](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324081908509.png)

- 强制覆盖不提示的方法：\cp

  \ ：表示使用原生命令执行 

##### 2.7 rm 删除文件或目录

- 基本语法

  rm [选项] deleFile   ---   递归删除目录中的所有内容

- 选项说明

  ![image-20250324083049670](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324083049670.png)

- 案例实操

  - 删除目录中的内容

    rm xiyou/mingjie/sunwukong.txt

  - 递归删除目录中所有内容

##### 2.8 mv 移动文件与目录或重命名

- 基本语法
  - mv oldNameFile newNameFile   ---   重命名
  - mv /temp/movefile /targetFolder   ---   移动文件
- 案例实操
  - 重名名 mv xiyou/dssz/sunwukong.txt  xiyou/dssz/houge.txt
  - 移动文件 mv xiyou/dssz/houge.txt ./

##### 2.9 cat 查看文件内容

查看文件内容，从第一行开始显示

- 基本语法

  cat [选项] 要查看的文件

- 选项说明

  ![image-20250324223710100](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324223710100.png)

- 经验技巧

  一般查看比较小的文件，一个屏幕就能显示全的

- 案例实操

  - 查看文件内容并显示行号

    cat -n houge.txt

##### 2.10 more 文件内容分屏查看器

​	more指令是一个基于VI编辑器的文本过滤器，它以全屏幕的方式按页显示文本文件的内容。more 指令中内置了若干快捷键，详见操作说明。

- 基本语法

  more 要查看的文件

  ![image-20250324224214961](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324224214961.png)

##### 2.11 less 分屏显示文件内容

​	less指令用来分屏查看文件内容，它的功能与more指令相似，但是比more指令更加强大，支持各种显示终端。less指令在现实文件内容时，并不是一次将整合文件加载之后才显示，而是根据显示需要加载内容，对于显示大型文件具有较高的效率

- 基本语法

  less 要查看的文件

- 操作说明

  ![image-20250324225320952](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324225320952.png)

##### 2.12 echo 

​	echo 输出内容到控制台

- 基本语法

  echo [选项] [输出内容]

  选项：

  -e：支持反斜线控制的字符转换

  ![image-20250324225531346](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250324225531346.png)

##### 2.13 > 输出重定向 >> 追加

- 基本语法
  - ls -l > 文件   ---   列表的内容写入文件 a.txt 中（覆盖写）
  - ls -al >> 文件 ---   列表的内容追加到文件 aa.txt 的末尾
  - cat 文件1 > 文件2   ---   将文件1的内容覆盖到文件2
  - echo "内容"   >>  文件   ---   将“内容”追加到文件中

##### 2.14 head 显示文件头部内容

​	head 用于显示文件的开头部分内容，默认情况下 head 指令显示文件的前10行内容

- 基本语法

  head 文件   ---   查看文件头10行内容

  head -n 5 文件   ---   查看文件头5行内容，5可以是任意行数

- 选项说明

  ![image-20250325075020204](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250325075020204.png)

##### 2.15 tail 输出文件尾部内容

​	tail 用于输出文件中尾部的内容，默认情况下tail指令显示文件的后10行内容

- 基本语法

  - tail 文件   ---   查看文件尾部10行内容
  - tail -n 5 文件   ---   查看文件尾部5行内容，5可以是任意行数
  - tail -f 文件   ---    实时追踪该文档的所有更新

- 选项说明

  ![image-20250325075414954](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250325075414954.png)

##### 2.16 ln 软链接

​	软链接也称为符号链接，类似于 windows 里面的快捷方式，有自己的数据块，主要存放了链接其他文件的路径。

- 基本语法 

  ls -s [源文件或目录] [软链接名]   ---   给原文件创建一个软链接

- 经验技巧

  删除软链接：rm -rf 软链接名，而不是 rm -rf r软链接名/

  如果使用 rm -rf 软链接名/ 删除，会把软链对应的真实目录下内容删掉

  查询：通过 ll 就可以查看，列表属性第1位是l，尾部会有位置指向。

##### 2.17 history 查看已经执行过的历史命令

- 基本语法

  history   ---   查看已经执行过的历史命令

#### 3. 时间日期类

- 基本语法 

  datt  [option]  [+format]

- 选项说明

  ![image-20250325082854682](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250325082854682.png)

- 参数说明

  ![image-20250325082927297](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250325082927297.png)

##### 3.1 date 显示当前时间

- 基本语法
  - date   ---   显示当前时间
  - date +%Y   ---   显示当前年份
  - date +%m   ---   显示当前月份
  - date +%d   ---   显示当前是那一天
  - date "+%Y-%m-%d %H:%M:%S"   ---   显示年月日时分秒

##### 3.2 显示非当前时间

date -d '1 days ago'   ---   显示昨天时间

date -d '-1 days ago'   ---   显示昨天时间

##### 3.3 date 设置系统时间

- 基本语法

  date -s 字符串时间

- 案例实操

  - 设置系统当前时间

    date -s "2017-06-19 20:52:18"

  - ntpdate + 服务器 ---  设置为当前时间

##### 3.4 cal 查看日历

- 基本语法

  cal [选项]   ---   不加选项，显示本月日历

- 选项说明

  ![image-20250326082432242](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250326082432242.png)

- 案例实操

  查看当前月的日历

  cal

- 查看2017年的日历

  cal 2017

#### 4.用户管理命令

##### 4.1 useradd 添加新用户

- 基本语法

  useradd 用户名   ---   添加新用户

  useradd -g 组名 用户名   ---   添加新用户到某个组

- 案例实操

  - 添加一个用户

    useradd ahu

    ll /home/ 

  - useradd -d /home/dave david   --- 创建新用户，并修改用户目录为dave

##### 4.2 passwd 设置用户密码

- 基本语法

  passwd 用户名   ---   设置用户密码

- 案例实操

  passwd ahu

##### 4.3 查看用户是否存在

- 基本语法

  id 用户名

- 案例实操

  id ahu

##### 4.4 cat /etc/passwd 查看创建了哪些用户

- 案例实操   

  cat /etc/passwd

##### 4.5 切换用户

- 案例实操

  su ahu

##### 4.6 查看登录用户信息

​	命令：who am i 

##### 4.7 sudo 设置普通用户具有root权限

- 添加 ahu 用户，并设置其密码

  ```sh
  useradd ahu
  passwd ahu
  ```

- 修改配置文件

  修改 /etc/sudoers 文件，找到下面一行（91行），在 root 下面添加一行，如下图所示：

  ```bash
  root ALL=(ALL)    ALL
  ahu ALL=(ALL)    ALL
  ```

  或者配置成采用 sudo 命令时，不需要输入密码：

  ```bash
  root ALL=(ALL)    ALL
  ahu ALL=(ALL)    NOPASSWD:ALL
  ```

  修改完毕，现在可以用 ahu 账号登录，然后用命令 sudo，即可获得 root 权限进行操作。

##### 4.8 userdel 删除用户

userdel devid

userdel -r david   ----  同时删除 /home 下对应的目录

##### 4.9 usermod 修改用户所属组

1. 基本用法

   usermod -g 用户组 用户名

2. 选项说明

   ![image-20250327080415459](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250327080415459.png)

#### 5.用户组管理命令

​	每个用户都有一个用户组，系统可以对一个用户组中的所有用户进行集中管理。不同 Linux 系统对用户组的规定有所不同。

​	如 Linux 下的用户属于与它同名的用户组，这个用户组在创建用户时同时创建。

​	用户组的管理设计用户组的添加，删除，修改。组的增加，删除和修改实际上是对 /etc/group 文件的更新。

##### 5.1 groupadd 新增组

1. 基本语法

   groupadd 组名

##### 5.2 groupdel 删除组

1. 基本语法 

   groupdel 组名

##### 5.3 修改组

1. 基本语法

   groupmod -n 旧组名 新组名

##### 5.4 cat /etc/group 查看创建了哪些组

#### 6.文件权限类

##### 6.1 文件属性

​	Linux 系统是一种典型的多用户系统，不同的用户处于不同的地位，拥有不同的权限。为了保护系统的安全性，Linux系统对不同的用户访问同一文件（包括目录文件）的权限做了不同的规定。在 Linux 中我们可以使用 ll 或 ls -l 命令来显示一个文件的属性以及文件属性的用户和组。

1. 从左到右的10个字符表示：

   ![image-20250327081805252](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250327081805252.png)

   1. 0 首位表示类型

      在 Linux 中第一个字符代表这个文件是目录，文件或链接文件等等。

       [ - ] 代表文件

      d 代表目录

      l  代表链接文档 （link file）

   2. 第1-3位确定属主（该文件的所有者）拥有该文件的权限。--- User

   3. 第4-6位确定属组（所有组的同组用户）拥有该文件的权限。--- Group

   4. 第7-9位确定其他用户拥有该文件的权限。 --- Other

2. rwx 作用文件和目录的不同解释

   1. 作用到文件：
   
      [r] 代表可读（read）：可以获取，查看
   
      [w] 代表可写（write）：可以修改，但是不代表可以删除该文件，删除一个文件的前提是对该文件所在的目录有写权限，才能删除该文件。
   
      [x] 代表可执行（execute）：可以被系统执行
   
   2. 作用到目录：
      [r] 代表可读（read）：可以读取，ls 查看目录内容
   
      [w] 代表可写（write）：可以修改，目录内创建+删除+重命名目录
   
      [x] 代表可执行（execute）：可以进入该目录
   
   3. 实操案例
   
      ![image-20250328082850442](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250328082850442.png)
   
      ![image-20250328082909353](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250328082909353.png)
   
      1.    如果查看的是文件：链接数指的是硬链接个数
      2.    如果查看的是文件夹：链接数指的是子文件个数

##### 6.2 chmod 改变权限

1. 基本语法

   ![image-20250328083453257](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250328083453257.png)

   第一种方法变更权限

   ​	chmod [{ugoa}{+-=}{rwx}]文件或目录 

   第二种方式变更权限

   ​	chmod [mode=421] [文件或目录]

2. 经验技巧

   u：所有者   g：所有组   o：其他人   a：所有人（u、g、o的总和）

   r=4 w=2 x=1   rwx=4+2+1=7
   
   修改整个文件夹里面所有文件的所有者，所属组，其他用户都具有可读可写可执行权限
   
   chmod -R 777 ahu/

##### 6.3 chown 改变所有者

1. 基本语法

   chown [选项] [最终用户] [文件或目录]   ---  改变文件或目录的所有者

2. 选项说明

   ![image-20250331083036841](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250331083036841.png)

##### 6.4 chgrp 改变所属组

1. 基本语法 [最终用户组] [文件或目录]   ---   改变文件或目录的所属组

   chgrp [最终用户组] [文件或目录]

#### 7.搜索查找类

##### 7.1 find 查找文件或者目录

​	find 指令将从指令目录向下递归地遍历其各个子目录，将满足条件的文件显示在终端。

1. 基本语法

   find [搜索范围] [选项]

2. 选项说明

   ![image-20250401080601227](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401080601227.png)

3. 案例实操

   1. 按文件名：根据名称查找/目录下的filename.cfg文件

      ![image-20250401081006001](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401081006001.png)

   2. 按拥有者：查找/home目录下，用户名称为ahu的文件

      ![image-20250401081314657](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401081314657.png)

   3. 按文件大小：在 /root 目录下查找大于2M的文件（+n 大于 -n小于 n等于）

      ![image-20250401081437575](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401081437575.png)

##### 7.2 locate 快速定位文件路径

​	locate 指令利用事先建立的系统重所有文件名称及路径的locate数据库实现快速定位给定文件。locate指令无需遍历整个文件系统，查询速度较快。为了保证查询结果的准确度，管理员必须定期更新locate时刻。

1. 基本语法

   locate 搜索文件

2. 由于 locate 指令基于数据库进行查询，所以第一次运行前，必须使用 updatedb 指令创建 locate 数据库。

3. 案例实操

   1. 查询文件夹

      ![image-20250401082420812](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401082420812.png)

   2. 查找命令所在位置的命令

      which [命令]

      whereis [命令]

##### 7.3 grep 过滤查找及“|” 管道符

​	管道符， “|” ，表示将前一个命令的处理结果输出传递给后面的命令处理。

1. 基本语法

   grep 选项 查找的内容 源文件

2. 选项说明

   ![image-20250401082804736](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401082804736.png)

3.  实际操作

   ![image-20250401083140970](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250401083140970.png)

#### 8.压缩和解压类

##### 8.1 gzip/gunzip 压缩

1. 基本语法

   gzip 文件   ---  压缩文件，只能将文件压缩为 *.gz 文件

   gunzip 文件.gz   ---   解压缩文件命令

2. 经验技巧

   1. 只能压缩文件不能压缩目录
   2. 不保留原来的文件
   3. 同时多个文件会产生多个压缩包

3. 案例实操

   ![image-20250403074343649](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403074343649.png)

##### 8.2 zip/unzip 压缩

1. 基本语法

   zip [选项] XXX.zip 将要压缩的内容   ---   压缩文件和目录的命令

   unzip [选项]   XXX.zip   ---   解压缩文件

2. 选项说明

   ![image-20250403074551901](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403074551901.png)

3. 经验技巧

   zip 压缩命令在 window/linux都通用，可以压缩目录且保留源文件。

4. 案例实操

   ![image-20250403075110489](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403075110489.png)

##### 8.3 tar 打包

1. 基本语法

   tar [选项] XXX.tar.gz 将要打包进去的内容   ---   打包目录，压缩后的文件格式 .tar.gz

2. 选项说明

   ![image-20250403075316998](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403075316998.png)

3. 案例实操

   ![image-20250403080022644](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403080022644.png)

#### 9. 磁盘查看和分区类

##### 9.1 du 查看文件和目录占用的磁盘空间

​	du: disk usage 磁盘占用情况

1. 基本语法

   du 目录/文件   ---   显示目录小每个子目录的此案使用情况

2. 选项说明

   ![image-20250403080227340](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403080227340.png)

3. 案例实操

   1. 查看当前用占用磁盘大小

      du -ch

##### 9.2 df 查看磁盘空间使用情况

​	df：disk free 空余磁盘

1. 基本语法

   df 选项   ---   列出文件系统的整体磁盘使用量，检查文件系统的磁盘空间占用情况

2. 选项说明

   ![image-20250403082550940](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403082550940.png)

3. 案例实操

   ![image-20250403083700006](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403083700006.png)

##### 9.3 lsblk 查看设备挂载情况

1. 基本语法

   lsblk   ---   查看设备挂载情况

2. 选项说明

   ![image-20250403083820124](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250403083820124.png)

##### 9.4 mount/umount 挂载/卸载

​	对于 Linux 用户来讲，不论有几个分区，分别分给哪一个目录使用，它总归就是一个根目录，一个独立且唯一的文件结构。

​	Linux中每个分区都是用来组成整个文件系统的一部分，它在用一种叫做“挂载”的处理方法，整个文件系统中包含了一整套的文件和目录，并将一个分区和一个目录联系起来，要载入的那个分区将使它的存储空间在这个目录下获得。

1. 挂载前准备（必须要有光盘或已经连接镜像文件）

2. 基本语法

   mount [-t vfstype] [-o options] device dir   ---   挂载设备

   umount 设备文件名或挂载点   ---   卸载设备

3. 参数说明

   ![image-20250408081747550](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250408081747550.png)

4. 案例实操

   ![image-20250408082308483](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250408082308483.png)

5. 设置开机自动挂载

   vim /etc/fstab

   ![image-20250408083024307](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250408083024307.png)

#####  9.5 fdisk 分区

1. 基本语法 

   fdisk -l   ---   查看磁盘分区详情

   ![image-20250408083543287](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250408083543287.png)

   fdisk 硬盘设备名   ---   对新增硬盘进行分区操作

2. 选项说明

   ![image-20250408083224667](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250408083224667.png)

3. 经验技巧

   该命令必须在root用户下才能使用

4. linux绑定磁盘功能说明

   1. 拥有一块新的磁盘
   2. fdisk /dev/sdb   ---   给sdb新磁盘创建分区
   3. mkfs -t xfs /dev/sdb1   ---   给sdb新磁盘的sdb1分区创建文件系统（xfs）
   4. mount /dev/sdb1 /home/guojinghu   ---   将该新磁盘的sdb1分区挂载到 /home/guojinghu 文件下

#### 10.进程管理类

##### 10.1 ps 查看当前系统进程状态

​	ps: process status 进程状态

1. 基本语法

   ps aux | grep xxx   ---   查看系统中所有进程

   ps -ef | grep xxx   ---   可以查看父子进程之间的关系

2. 选项说明

   ![image-20250410075319973](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250410075319973.png)

3. 功能说明

   1. ps aux 显示信息说明

      USER：该进程是由哪个用户产生的

      PID：进程的ID号

      %CPU：该进程占用CPU资源的百分比，占用越高，进程越耗费资源

      %MEM：该进程占用物理内存的百分比，占用越高，进程越耗费资源

      VSZ：该进程占用虚拟内存的大小，单位KB

      RSS：该进程占用实际物理内存的大小，单位KB

      TTY：该进程是在哪个终端中运行，对于Centos来说，tty1是图形化终端，tty2-tty6是本地的字符界面终端，pts/0-255代表虚拟终端。

      STAT：进程状态。常见的状态有：R：运行状态 S：睡眠状态 T：暂停状态 Z：僵尸状态 s：包含子进程 l：多线程 +：前台显示

      START：该进程的启动时间

   2. ps -ef 显示信息说明

      UID：用户ID

      PID：进程ID

      PPID：父进程ID

      C：CPU用于计算执行优先级的因子。数值越大，表明进程是CPU密集型运算。执行优先级会降低，数值越小，表明进程是I/O密集型运算，执行优先级会提高。

      STIME：进程启动的时间

      TTY：完整的终端名称

      TIME：CPU时间

      CMD：启动进程所用的命令和参数

4. 经验技巧

   - 如果想查看进程的CPU占用率和内存占用率，就可以使用 aux
   - 如果想查看进程的父进程ID可以使用 ef

##### 10.2 kill 终止进程

1. 基本语法

   kill [选项] 进程号   ---   通过进程号杀死进程

   Killall 进程名称   ---   通过进程名称杀死进程，也支持通配符，这在系统因负载过大而变得慢时很有用。

2. 选项说明

   ![image-20250410083216666](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250410083216666.png)

3. 实操案例

   1. kill -9 5102
   2. killall firefox


##### 10.3 pstree 查看进程树

1. 基本语法

   pstree [选项]

2. 选项说明

   ![image-20250411080510761](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250411080510761.png)

3. 实操案例

   1. 显示进程pid:  pstree -p
   2. 显示进程所属用户：pstree -u

##### 10.4 top 实时监控系统进程状态

1. 基本命令

   top [选项]

2. 选项说明

   ![image-20250411081220483](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250411081220483.png)

##### 10.5 netstat 显示网络状态和端口号占用信息

1. 基本语法

   netstat -anp | grep 进程号   ---   查看该进程网络信息

   netstat -nlp | grep 端口号   ---   查看网络端口号占用情况

2. 选项说明

   ![image-20250411083139651](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250411083139651.png)

3. 实际操作

   ![image-20250411083249800](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250411083249800.png)

#### 11 crontab 系统定时任务

##### 11.1 crontab 服务管理

1. 重新启动 ctond 服务

##### 11.2 crontab 定时任务复制

1. 基本语法

   crontab [选项]

2. 选项说明

   ![image-20250415082651803](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250415082651803.png)

3. 参数说明

   crontab -e

   1. 进入 crontab 编辑界面。会打开 vim 编辑你的工作。

      “* * * * * 执行的任务”

      ![image-20250415083114190](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250415083114190.png)
   
   2. 特殊符号
   
      ![ ](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250416075815218.png)
   
   3. 特定时间执行命令
   
      ![image-20250416075847526](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250416075847526.png)

## 软件包管理

#### 1. RPM

##### 1.1 RPM 概述

​	RPM（RedHat Package Manager），RedHat 软件包管理工具，类似于windows里面的setup.exe。是Linux这系列操作系统里面的打包工具，它虽然是RedHat的标志，但是理念是通用的。

​	RPM包的名称格式

​	Apache-1.3.23-11.i386.rpm

​		"apache" 软件名称		

​		"1.3.23-11" 软件的版本号，主版本和此版本

​		"i386" 是软件所运行的硬件平台，Intel 32位处理器的统称

​		"rpm" 文件扩展名。代表 RPM 包

##### 1.2 RPM 查询命令（rpm -qa）

1. 基本语法

   rpm -qa   ---   查询所安装的所有 rpm 软件包

2. 经验技巧

   由于软件包比较多，一般都会采取过滤。 rpm -qa | grep rpm软件包

3. 案例实操

   1. 查询 firefox软件安装情况

      rpm -qa | grep firefox

      ![image-20250416081608344](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250416081608344.png)

##### 1.3 RPM 卸载命令（rpm -e）

1. 基本语法

   1. rpm -e RPM软件包
   2. rpm -e --nodeps 软件包

2. 选项说明

   ![image-20250416081928605](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250416081928605.png)

3. 案例实操

   1. 卸载 firefox 软件
      1. rpm -e firefox

##### 1.4 RPM 安装命令（rpm -ivh）

1. 基本语法

   rpm -ivh RPM包全名

2. 选项说明

   ![image-20250416082306488](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250416082306488.png)

#### 2. YUM 仓库配置

##### 2.1 YUM 概述

​	YUM（全称为 Yellow dog Updater, Modified）是一个在 Fedora 和 RedHat 以及 CentOS 中的 Shell 前端软件包管理器。基于 RPM 包管理，能够从指定的服务器自动下载 RPM 包并且安装，可以自动处理依赖性关系，并且一次安装所有依赖的软件包，无须繁琐低一次次下载，安装。

![image-20250417075058583](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417075058583.png)

##### 2.2 YUM 的常用语法

1. 基本语法

   yum [选项] [参数]

2. 选项说明

   ![image-20250417075209216](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417075209216.png)

3. 参数说明

   ![image-20250417075226039](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417075226039.png)

4. 案例实操

   1. 采用 yum 方式安装 firefox

      yum -y install firefox

##### 2.3 修改网络 YUM 源

​	默认的系统 YUM 源，需要连接国外 apache 网站，网速比较慢，可以修改关联的网络 YUM 源为国内镜像的网站，比如 163，aliyun等。

1. 安装 wget, wget 用来从指定的 URL 下载文件

   yum install wget

2. 在 /etc/yum.repos.d/ 目录下，备份默认的 repos 文件

   ![image-20250417081016190](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417081016190.png)

3. 下载网易163 或者 aliyun 的 repos 文件，任选其一

   ![](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417081102684.png)

4. 使用下载好的 repos 文件替换默认的 repos 文件

   例如：用 CentOS7-Base-163.repo 替换 Centos-Base.repo

   ![image-20250417081533135](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417081533135.png)

5. 清理旧缓存数据，缓存新数据

   yum clean all

   yum makecache

6. 测试

   yum list | grep firefox

   yum -y install firefox

## 克隆虚拟机

#### 1. 克隆

1. 从现有虚拟机（关机状态）克隆出新虚拟机，邮件选择管理 => 克隆

2. 指定好虚拟机名称，本机存储地址之后点击创建

3. 启动之后修改虚拟机ip地址 （/etc）

   ![image-20250417083331463](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417083331463.png)

   ![image-20250417083404440](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417083404440.png)

## Shell

### 1. Shell 概述

​	Shell 是一个命令行解释器，它接收应用程序/用户命令。然后调用操作系统内核。

​		 ![image-20250417083735976](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250417083735976.png)

​	Shell 还是一个功能相当强大的编程语言，易编写，易调试，灵活性强。

1. Linux 提供的 Shell 解析器

   ![image-20250418074447857](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250418074447857.png)

2. bash 和 sh 的关系

   ![image-20250418074752897](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250418074752897.png)

### 2. Shell 脚本入门

1. 脚本格式

   脚本以 #!/bin/bash开头（指定解释器）

2. 第一个Shell脚本：helloworld.sh

   1. 需求：创建一个shell脚本，输出 helloworld

   2. 案例实操

      ![image-20250418075437481](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250418075437481.png)

      在 hello.sh 文档中输入如下内容：
      ![image-20250418075612373](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250418075612373.png)

   3. 脚本的常用执行方式

      1. 第一种：采用 bash 或 sh+ 脚本的相对路径或决对路径（不用赋予脚本+x权限）
      2. 第二种：采用输入脚本的绝对路径或相对路径执行脚本（必须具有可执行权限+x）
      3. 第三种：在脚本的路径前加上 “.” 或者 source

      原因：

      ​	前两种方式都是在当前shell中打开一个shell来执行脚本内容。当脚本内容结束，则子shell关闭，回到父shell中。

      ​	第三种，也就是使用在脚本路径前加 “.” 或者 source 的方式，可以使脚本内容在当前 shell 里执行，而无需打开子 shell！这也是为什么我们每次要修改完/etc/profile文件后，需要source一下的原因。

      ​	开子shell与不开子shell的区别就在于，环境变量的继承关系，如在子shell中设置的当前变量，父shell是不可见的。

### 3. 变量

#### 3.1 系统预定义变量

1. 常用系统变量

   $HOME   $PWD   $SHELL   $USER等

2. 案例实操

   1. 查看系统变量的值

      echo $HOME

   2. 显示当前 Shell 中所有变量 set

      ![image-20250418081924977](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250418081924977.png)

#### 3.2 自定义变量

 

1. 1. 基本语法

      1. 定义变量：变量名=变量值，注意=后面不能有空格
      2. 撤销变量：unset 变量，名
      3. 声明静态变量：readonly 变量，注意：不能unset

   2. 变量定义规则

      1. 变量名称可以由字母，数字和下划线组成，但是不能以数字开头，**环境变量名建议大写**。
      2. 等号两侧不能有空格
      3. 在 bash 中，变量默认类型都是字符串类型，无法直接进行数值运算
      4. 变量的值如果有空格，需要使用双引号或单引号括起来。

   3. 案例实操

      ![image-20250422081056361](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250422081056361.png)

      ![image-20250422081125698](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250422081125698.png)

#### 3.3 特殊变量

##### 3.3.1 $n

1. 基本语法

   $n   ---   功能描述：n为数字，$0代表该脚本名称，$1-$9代表第一到第九个参数，十以上的参数需要用大括号包含，比如${10}

2. 案例实操

   ![image-20250422081902447](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250422081902447.png)

##### 3.3.2 $#

1. 基本语法

   $#   ---   功能描述：获取所有输入参数个数，常用于循环，判断参数的个数是否正确以及加强脚本的健壮性。

2. 实操案例

   ![image-20250422082859138](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250422082859138.png)

##### 3.3.3 $* 、$@

1. 基本语法

   $* --- 功能描述：这个变量代表命令行中所有的参数，$*把所有的参数看成一个整体。

   $@---功能描述：这个变量也代表着命令行中所有参数，不过$@把每个参数区分对待

2. 实操案例

   ![](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250422083418366.png)

##### 3.3.4 $?

1. 基本语法

   $?   ---   功能描述：最后一次执行的命令的返回状态。如果这个变量的值为0，证明上一个命令正确执行；如果这个变量的值为非0（具体是哪个数，有命令自己来决定），则证明上一个命令执行不正确了。

2. 案例实操

   判断 helloworld.sh 脚本是否正确执行？

   ![image-20250422083730383](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250422083730383.png)

### 4. 运算符

1. 基本语法

   "$((运算式))" 或 "$[运算式]"

2. 案例实操

   计算 (2+3) * 4的值

   s=$[(2+3)*4]

   echo $s

### 5. 条件判断

1. 基本语法

   1. test condition

   2. [ condition ]  注意 condition 前后要有空格

      注意：条件非空即为true。[atguigu]返回true，[ ]返回false

 



















 











