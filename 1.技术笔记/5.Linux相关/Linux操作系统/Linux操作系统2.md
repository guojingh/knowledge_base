#### 10.2 awk

​	一个强大的文本分析工具，把文件逐行读入，以空格为默认的分隔符将每行切片，切开后的部分再进行分析处理。

1. 基本用法

   awk [选项参数] '/pattern1/{action1} /pattern2/{action2}...' filename

   Pattern: 表示 awk 在数据中查找的内容，就是匹配模式

   action: 在找到匹配内容时所执行的一系列命令

2. 选项参数说明

   ![image-20250615113519044](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250615113519044.png)

3. 案例实操

   1. 搜索passwd文件以root关键字开头的所有行，并输出该行的第7列

      ```shell
      v_guojinghu@MacBook-Pro-4 ~ % cat /etc/passwd | awk -F ":" '/^root/ {print $7}' 
      /bin/sh
      ```

   2. 搜索 passwd 文件以root关键字开头的所有行，并输出该行的第一列和第7列，中间以‘，’号分隔。

      ```sh
      v_guojinghu@MacBook-Pro-4 ~ % cat /etc/passwd | awk -F ":" '/^root/ {print $1","$7}' 
      root,/bin/sh
      ```

   3. 只显示/etc/passwd的第一列和第七列，以逗号分隔，且在所有行前面添加列名user,shell在最后一列添加“dahaige, /bin/zuishuai”

      ```sh
      v_guojinghu@MacBook-Pro-4 ~ % cat /etc/passwd | awk -F ":" 'BEGIN{print "user, shell"}{print $1""$7} END{print "end of file"}' 
      user, shell
      ```

   4. 将 passwd 文件中的用户id增加数值1并输出

      ```sh
      v_guojinghu@MacBook-Pro-4 ~ % cat /etc/passwd | awk -F ":" '{print $3+1}' 
      ```

   5. 将 passwd 文件中的用户id增加数值（参数版）并输出

      ```sh
      v_guojinghu@MacBook-Pro-4 ~ % cat /etc/passwd | awk -v i=2 -F ":" '{print $3+i}' 
      ```

   6. awk的内置变量

      ![image-20250615115900447](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250615115900447.png)

   7. 案例实操

      1. 统计 passed 文件名，每行的行号，每列的列数

         ``` sh
          v_guojinghu@MacBook-Pro-4 ~ % cat /etc/passwd | awk -F ":" '{print "文件名： "FILENAME" 行号："NR" 列数："NF""}' 
         ```

      2. 查询 ifconfig 命令输出结果中的空行所在的行号

         ```sh
         v_guojinghu@MacBook-Pro-4 ~ % ifconfig| awk '/^$/ {print NR}'
         ```

      3. 切割IP

         ```sh
         ifconfig | awk '/netmask/ {print $2}' 
         ```

### 11. 综合应用案例

#### 11.2 发送消息

​	我们利用 Linux 自带的 mesg 和 write 工具，向其它用户发送消息。

​	需求：实现一个向某个用户快速发送消息的脚本，输入用户名作为第一个参数，后面直接跟要发送的消息。脚本需要检测用户是否登录在系统中，是否打开消息功能，以及当前发送消息是否为空。

​	脚本实现如下：

```shell
#!/bin/bash

# 查看用户是否登录
login_user=$(who | grep -i -m 1 $1 | awk '{print  $1}')

if [ -z $login_user ]
then
	echo "$1 不在线！"
	echo "脚本退出.."
	exit
fi

is_allow=$(who -T | grep -i -m 1 $1 | awk '{print $2}')

is [ $is_allow != "+" ]
then
	echo "$1 没有开启消息功能"
	echo "脚本退出"
	exit
fi

# 确认是否有消息发送
if [ -z $2 ]
then	
	echo "没有消息发送"
	echo "脚本退出"
	exit
fi

# 从参数中获取要发送的消息
whole_msg=$(echo $* | cut -d " " -f 2-)

# 获取用户登录的终端
user_terminal=$(who | grep -i -m 1 $1 | awk '{print $2}')

# 写入要发送的消息
echo $whole_msg | write $login_user $user_terminal

if [ $? != 0 ]
then 
	echo "发送失败"
else
	echo "发送成功"
fi

```

