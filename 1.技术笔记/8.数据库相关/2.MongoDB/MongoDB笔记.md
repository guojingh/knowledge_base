# MongoDB

## 1.简介

### 1.1 什么是MongoDB?

MongoDB 是面向文档的NoSQL数据库，用于大量数据存储。MongoDB是一个在2000年代中期问世的数据库。属于NoSQL数据库的类别。

### 1.2 什么是NoSQL数据库？

NoSQL数据库即 No Only SQL 的缩写，即“不仅仅是SQL”。区别于关系型数据库，它们不保证关系型数据库的ACID特性。当然，它们的功能都是存储数据。

### 1.3 MongoDB 功能？

1. 每个数据库度包含集合，而集合又包含文档。每个文档可以具有不同数量的字段。每个文档的大小和内容可以互不相同。
2. 文档结构更符合开发人员如何使用各自的编程语言构造其类和对象。开发人员经常会说他们们的类不是行和列，而是具有健值对的清晰结构。
3. 从NoSQL数据库的简介中可以看出，行（或在MongoDB中调用的文档）不需要预先定义架构。相反，可以动态创建字段。
4. MongoDB中可用的数据模型使我们可用轻松地表示层次结构关系，存储数组和其他更复杂的结构。
5. 可伸缩性-MongoDB环境具有很高的可伸缩性。全球各地的公司已经定义了自己的集群，其中一些集群运行着100多个节点，数据库中包含大约百万个文档。

### 1.4 MongoDB 架构的关键组件

下面是MongoDB中使用的一些常用术语

1. _id: 这是每个MongoDB文档中必填的字段。表示MongoDB文档中的唯一值。类似于文档的主键，如果创建的新文档中没有 _id字段，MongoDB 将自动创建该字段即一个24位唯一标识符。
2. 集合：这是MongoDB文档的分组。集合等效于在其它关系型数据库中创建的表。集合存在于单个数据库中。从介绍中看出，集合不强制执行任何结构。
3. 游标：这是指向查询结构集的指针。客户可以遍历游标以检索结果。
4. 数据库：集合的容器，每个数据库在文件系统上都有自己的文件集。MongoDB服务器可以存储多个数据库。
5. 文档：MongoDB集合中的记录基本上称为文档。文档包含字段名称和值。
6. 字段：文档中的名称/值对。一个文档具有零个或多个字段。字段类似于关系型数据库中的列。

### 1.5 为什么要使用MongoDB?

1. 面向文档：由于MongoDB是NoSQL数据库，他不是以关系类型的格式存储数据，而是将数据存储在文档中。这使得MongoDB非常灵活，可以适应实际的业务环境和需求。
2. 临时查询：支持按字段，范围查询和正则表达式搜索。可以查询返回文档中的特定字段。
3. 索引：可以创建索引提高搜索性能，MongoDB文档中任何字段都可以建立索引。
4. 复制：提供副本集提高可用性。
5. 负载均衡：使用分片的概念，通过在多个MongoDB实例之间拆分数据来水平拓展。

### 1.6 MongoDB和RDBMS之间的区别

![image-20250314111159746](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314111159746.png)

## 2.NoSQL

### 2.1 什么是NoSQL?

NoSQL是一种非关系型DMS，不需要固定的架构，可以避免joins链接，并且易于扩展。NoSQL数据库用于庞大数据存储需求的分布式数据存储。NoSQL用于大数据和实时Web应用程序。例如，像Twitter，facebook，Google这样大型公司，每天可能产生TB级别的用户数据。

传统的RDBMS使用SQL语法来存储和查询数据。相反，N oSQL数据库系统包含可存储结构化，半结构化，非结构化和多态数据的多种数据库技术。

![image-20250314112223216](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314112223216.png)

### 2.2 为什么使用NoSQL？

NoSQL数据库的概念在处理大量数据的互联网巨头（例如：Google,Facebook,Amazon等）中变得很流行，使用RDBMS处理海量数据时，系统响应时间慢。

为了解决此问题，当然可以通过升级现有硬件来“横向扩展”我们的系统，但是这个成本很高。

这个问题的代替方案是在负载增加时将数据库负载分配到多个主机上。这种方法称为横向扩展。

![image-20250314112940119](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314112940119.png)

NoSQL数据库时非关系型数据库，因此在设计时考虑到we b应用程序，比关系型数据库更好地扩展。

### 2.3 NoSQL数据库的简要历史

- 1998年-Carlo Strozzi在他的轻量级开源关系数据库中使用术语NoSQL
- 2000-图形数据库Neo4j启动
- 2004年-推出Google BigTable
- 2005年-启动CouchDB
- 2007年-发布有关Amazon Dynamo的研究论文
- 2008年-Facebook开源Cassandra项目
- 2009年-重新引入NoSQL术语

### 2.4 NoSQL功能

#### 2.4.1 非关系型

- NoSQL数据库从不遵循关系模型
- 切勿为tables 提供固定的记录
- 使用自包含的聚合或BLOB
- 不需要对象关系映射和数据规范化
- 没有复杂的功能，例如查询语言，查询计划者

#### 2.4.2 动态架构

NoSQL数据库是无模式的或具有宽松模式的数据库

不需要对数据架构进行任何形式的定义

提供同一域中的异构数据结构

![image-20250314113648390](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314113648390.png)

#### 2.4.3 简单的API

- 提供易于使用的界面，用于存储和查询提供的数据
- API允许进行低级数据操作和选择方法
- 基于文本的协议，通常与带有JSON的HTTP REST一起使用
- 多数不使用基于标准的查询语言
- 支持Web的数据库作为面向互联网的服务运行

#### 2.4.4 分布式

- 可以以分布式方式执行多个NoSQL数据库
- 提供自动缩放和故障转移功能
- 通常可牺牲ACID概念来实现可伸缩性和吞吐量
- 分布式节点之间几乎没有同步复制，多为异步多主复制，对等，HDFS复制
- 仅提供最终的一致性
- 无共享架构。这样可以减少协调并提高分布。

#### 2.4.5 NoSQL 不共享

![image-20250314114718985](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314114718985.png)

### 2.5 NoSQL 数据库的类型

![image-20250314140723374](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314140723374.png)

NoSQL 数据库主要有四类。这些类别中的每一个都有其独特的属性和局限性。没有特定的数据库可以完美地解决所有问题。应该根据产品选择一个数据库。

- 基于健值对
- Column-oriented Graph
- Graphs-based
- 面向文档

#### 2.5.1 基于健值对

数据存储在健/值对中。它以这种方式设计可以处理大量数据和繁重的工作。

健值对存储数据库将数据存储位哈希表，其中每一个健都是唯一的，并且值可以是JSON,BLOB(二进制大对象)，字符串等。

![image-20250314141326002](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314141326002.png)

它是NoSQL数据库等最基本类型之一。这种NoSQL数据库用作集合，dictionaries，关联数组等。健值存储可帮助开发人员存储较少架构等数据。它们最合适购物车中的物品。

Redis，Dynamo，Riak是健值存储数据库的一些示例。它们全部基于亚马逊的 [Dynamo](https://arthurchiao.art/blog/amazon-dynamo-zh/)论文。

#### 2.5.2 基于列

面向列的数据库在列上工作，基于 Google 的 [BigTable](https://arthurchiao.art/blog/google-bigtable-zh/) 论文。每列都单独处理。单列数据库的值连续存储。

![image-20250314142819127](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314142819127.png)

在聚合查询（例如SUM,COUNT,AVG,MIN等）上提供了高性能，因为数据在列中随时可用。基于列的NoSQL数据库被广泛用于管理数据仓库，商业智能，CRM，图书馆书籍目录，Base，Cassandra，HBase，Hypertable是基于列的数据库。

#### 2.5.3 面向文档

面向文档的 NoSQL DB 将数据存储和检索为健值对，但值部分存储为文档。该文档以JSON或XML格式存储。DB可以理解为该值，并且可以查询该值。

![image-20250314143709331](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314143709331.png)

在上图中的左侧，我们可以看到有行和列，而在右侧，我们有一个文档数据库，该数据库的结构与Json类似。现在，对于关系数据库，必须知道拥有哪些列，依此类推。但是，对于文档数据库，具有JSON对象之类的数据存储。我们不需要定义，以便使其灵活。

文档类型主要用于CMS系统，博客系统，实时分析和电子商务应用程序。它不应该用于需要多种操作或针对不同聚合结构进行查询的复杂交易。

Amazon SimpleDB，CouchDB，MongoDB，Riak，Lotus Notes，MongoDB是流行的基于文档的DBMS系统。

#### 2.5.4 Graph-Based (基于图形)

图类型数据库存储实体以及这些实体之间的关系。实际存储为节点，关系作为边。一条边给出了节点之间的关系。每个节点和边缘都有唯一的标识符。

![image-20250314144408004](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314144408004.png)

与关系型数据库中表的松散连接相比，Graph 数据库本质上是多关系的。遍历关系很快，因为它们已经被存在于数据库中，因此无需计算它们。

基于图形数据库主要用于社交网络，物流，空间数据。

Neo4J, Infinite Graph, OrientDB, FlockDB是一些比较流行的基于图形的数据库。

### 2.6 NoSQL的查询机制工具

最常见的数据检索机制是基于 REST 的值检索，该值基于其key/ID，获取数据。

文档存储数据库提供了更加困难的查询，因为它们了解健值对中的值。例如，CouchDB允许使用MapReduce定义视图

### 2.7 什么是CAP定理

分布式数据存储不可能同时满足CAP，只能满足CAP其中的两部分。

- 一致性（**C**onsistency） （等同于所有节点访问同一份最新的数据副本）
- [可用性](https://zh.wikipedia.org/wiki/可用性)（**A**vailability）（每次请求都能获取到非错的响应——但是不保证获取的数据为最新数据）
- [分区容错性](https://zh.wikipedia.org/w/index.php?title=网络分区&action=edit&redlink=1)（**P**artition tolerance）（以实际效果而言，分区相当于对通信的时限要求。系统如果不能在时限内达成数据一致性，就意味着发生了分区的情况，必须就当前操作在C和A之间做出选择[[3\]](https://zh.wikipedia.org/wiki/CAP定理#cite_note-3)。）

### 2.8 NoSQL的优势

- 可用作主要数据库或分析数据源
- 大数据能力
- 没有单点故障
- 轻松复制
- 无需单独的缓存层
- 它提供了快速的性能和水平可扩展性
- 可以平等地处理结构化，半结构化和非结构化数据
- 易于使用且灵活的面向对象编程
- NoSQL数据库不需要专用的高性能数据库
- 支持关键的开发语言和平台
- 比使用RDBMS易于实现
- 它可以用作在线应用程序的主要数据源
- 处理大数据，以管理数据的速度，多样性，数量和复杂性
- 在分布式数据库和多数据中心操作方面表现出色
- 无需使用特定的缓存层子存储数据
- 提供灵活的架构设计，可以轻松进行更改而不会造成停机或者服务中断

### 2.9 NoSQL缺点

- 没有标准化规则
- 有限的查询功能
- RDBMS数据库工具和工具相对成熟
- 它不提供任何传统的数据库功能，例如同时执行多个事务的一致性
- 当数据量增加时，由于密钥变得困难，很难维护唯一值
- 对于新开发者而言，学习曲线是僵硬的
- 开源选项在企业中并不受欢迎

## 3. Docker 快速安装MongoDB

### 3.1 拉取 MongoDB 镜像

```sh
docker pull mongo
```

### 3.2 创建文件夹

```sh
mkdir -p /home/mongo/conf/
mkdir -p /home/mongo/data/
mkdir -p /home/mongo/logs/
```

### 3.3 新增mongod.conf文件

```sh
cd /home/mongo/conf && vi mongod.conf
```

mongod.conf文件内容：

```yaml
# 数据库文件存储位置
dbpath = /data/db
# log文件存储位置
logpath = /data/log/mongod.log
# 使用追加的方式写日志
logappend = true
# 是否以守护进程方式运行
# fork = true
# 全部ip可以访问
bind_ip = 0.0.0.0
# 端口号
port = 27017
# 是否启用认证
auth = true
# 设置oplog的大小(MB)
oplogSize=2048
```

### 3.4 新增mongod.log文件

```sh
cd /home/mongo/logs/ && vi mongod.log

##log文件不需要内容##
chmod  777 mongod.log 
```

### 3.5 docker 容器构建以及启动mongodb

```sh
cd /
docker run -it \
	--name mongodb \
	--restart=always \
    --privileged \
    -p 27017:27017 \
    -v /home/mongo/data:/data/db \
    -v /home/mongo/conf:/data/configdb \
    -v /home/mongo/logs:/data/log/  \
    -d mongo:latest \
    -f /data/configdb/mongod.conf
```

### 3.6 进入容器创建账号密码

```sh
##进入容器##
docker exec -it mongodb /bin/bash

##进入mongodb shell##
mongo

##切换到admin库##
> use admin

##创建账号/密码##
db.createUser({ user: 'admin', pwd: 'admin', roles: [ { role: "userAdminAnyDatabase", db: "admin" } ] });
```

## 4. MongoDB 具体操作

### 4.1 创建数据库和集合

在MongoDB中，第一步也是创建数据库和集合。数据库用于存储所有集合，而集合又用于存储所有文档。这些文档将依次包含相关的“字段名”和“字段值”。

下面的图展示了文档结构的例子。

![image-20250314164039965](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250314164039965.png)

#### 4.1.1 使用 ‘use’ 命令创建数据库

在MongoDB中创建数据库：

```sql
use EmployeeDB
```

在“use”命令创建MongoDB中的数据库。如果数据库不存在，将创建一个新的数据库。创建成功后，MongoDB将自动切换到创建的数据库。

#### 4.1.2 使用 insert() 创建集合/表

创建集合的最简单的方法是一条记录（不过是由字段名称和值组成的文档）插入到集合中。如果该集合不存在，则会创建一个新的集合。

**注：在MongoDB中，用户于特定数据库是相关联的。这意味着每个用户是在某个特定的数据库上创建的，并且其权限也是基于该数据库的。**

**赋予一个用户其他数据库的权限：**

```sql
// 切换到用户所在数据库
use testDB

// 更新用户角色
db.updateUser("testUser", 
	{
		roles: [
			{role: "readWrite", db: "testDB"},  // 原有的权限
			{role: "readWrite", db: "reportsDB"}  // 新增的权限
		]
	}
)
```

**创建只有一条记录的集合**

```sql
db.Employee.insert (
	{
		"雇员编号": 1,
		"EmployeeName": "马丁"
	}
)
```

#### 4.1.3 使用 insert() 命令添加文档

MongoDB 提供 insert() 命令将文档插入到集合中。下面的例子是如何完成此操作

```sql
db.Employee.insert (
	{
		"Employeeid": 1
		"EmployeeName": "Smith"
	}
)
```

1. 该命令第一部分是 “insert statement”，它就是用于将文档插入集合中的语句。
2. 该命令第二部分是添加字段名称和字段值，换句话说，集合中要包含的文档是什么内容。

#### 4.1.4 使用 insert() 在 MongoDB 中插入数组

“insert” 命令也可以一次将多个文档插入到集合中。下面我们操作如何一次插入多个文档。

1. 创建一个名为myEmployee的JavaScript变量来保存文档数据
2. 将具有字段名称和值的所需文档添加到变量
3. 使用insert命令来将文档数组插入集合中

```javascript
var myEmployee = [
  {
    "EmployeeId": 1,
    "EmployeeName": "Smith"
  },
  {
    "EmployeeId": 2,
    "EmployeeName": "Mohan"
  },
  {
    "EmployeeId": 3,
    "EmployeeName": "Joe"
  }
]

// 将文档插入到集合中
db.Employee.insert(myEmployee);
```

#### 4.1.5 以Json格式打印

使用Json格式查看输出：

```
db.Employee.find().forEach(printjson);
```

![image-20250317143001280](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317143001280.png)

**注：上述命令可能需要在mongosh 上执行，在有些数据库连接工具上面可能执行不成功。**

### 4.2 主键

#### 4.2.1 MongoDB 中的主键是什么？

在MongoDB中，_id 字段是集合的主键，以便可以在结合中唯一地标识每个文档。_id字段包含唯一地ObjectID值。

默认情况下，在集合插入文档时，如果没有在字段名称中添加带有id的字段，则MongoDB会自动添加一个Object id 字段。

![image-20250317143001280](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317143001280.png)

上面查询到的集合中的文档，可以看到该集合中每个文档的ObjectI值。

如果确保在创建集合时MongoDB不会创建id字段，并且要制定自己的id作为集合的_id，则需要在创建集合时明确定义它。

例子如下：

```sql
db.Employee.insert({_id: 10, "EmployeeName": "Smith"})
```

### 4.3 查询

#### 4.3.1 基本查询

```sql
db.Employee.find.forEach(printjson);
```

输出显示集合中存在的所有文档，我们还可以向查询中添加条件，以便我们可以根据特定条件获取文档。

1. 在集合中查找名称为“Smith”的Employee，因此我们将过滤条件输入为 EmployeeName:"Smith"

   ![image-20250317145429829](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317145429829.png)

2. 搜索哪些字段值大于指定值的文档，代码示例

   ![image-20250317145726133](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317145726133.png)

​	返回所有员工ID大于2的文档。

#### 4.3.2 游标教程

当使用 db.collection.find() 函数在集合中搜索文档时，结果将返回指向文档集合的指针，该指针称为游标。

默认情况下，返回查询结果时，游标将自动进行迭代。当然可以一个接一个明确展示游标中返回的结果目录。下面的例子中，告诉我们如何完成此操作。

```js
var myEmployee = db.Employee.find({EmployeeId: {$gt:2}});

while(myEmployee.hasNext())
  {
    print(tojson(myEmployee.next()));
  }
```

![image-20250317151054186](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317151054186.png)

**代码说明：**

1. 将查询到的结果集即ID大于2的Employee并将其赋值给JavaScript变量---“myEmployee”
2. 接下来，使用 while 循环遍历作为查询一个部分返回所有的文档
3. 最后，对于每个文档，我们以Json格式输出该文档的详细信息。

#### 4.3.3 使用 limit 查询结果

**什么是查询修饰符？**

MongoDB 提供了查询修饰符，例如：‘limit’ 和 ‘Orders’ 字句，以在执行查询时提供更大的灵活性。

此修饰符用于限制查询结果集中返回的文档数。下面的例子展示如何完成此操作。

```sql
db.Employee.find().limit(2).forEach(printjson);
```

**使用Limit子句将要返回的文档数限制为2**

![image-20250317151947570](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317151947570.png)

#### 4.3.4 降序排序

可以根据集合中任何健的升序排序和讲叙指定要返回的文档的顺序。

```sql
db.Employee.find().sort({EmployeeId:-1}).forEach(printjson)
```

**代码说明：**

上述代码采用sort函数，该函数返回集合中的所有文档，然后使用修饰符更改返回记录的顺序，这里的-1表示要根据EmployeeId的降序返回文档。

![image-20250317152437786](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317152437786.png)

上述图片显示了按EmployeeId降序返回的文档。

值为1则按照升序。

#### 4.4.4 Count() 函数

聚合的概念是对返回的结果进行计算。例如，假设我们想知道根据触发的查询集合中文档的数量是多少，那么MongoDB提供了 count() 函数。

```sql
db.Employee.countDocuments()
```

**代码说明**

上述代码执行计算功能。

![image-20250317154955514](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250317154955514.png)

### 4.5 remove() 删除文档

在 MongoDB 中，db.collection.remove() 方法用于从集合中删除文档。所有文档都可以从集合中删除，也可以仅从符合特定条件的文档中删除。

如果仅发出remove命令，则所有文档将从集合中删除。

下面的代码示例演示如何从集合中删除特定文档。

```
db.Employee.remove({EmployeeId:22})
```

### 4.6 update() 更新文档

MongoDB提供了 update() 命令来更新集合的文档。仅仅更新要更新的文档，可以将条件添加到新更新语句，以便更新选定的文档。

该命令中的基本参数就是需要更新文档的条件，其次是需要执行的修改。

```sql
db.Employee.update(
	{"EmployeeId": 1},
	{$set: {"EmployeeName": "NewMartin"}}
);
```

**更新多个值**

1. 发出更新命令

2. 选择要用于确定需要更新哪个文档的条件。在实例中，我们希望更新员工ID为 “1” 的文档；

3. 选择要修改的字段名称，并相应地输出其新值

   ```sql
   db.Employee.update(
   	{"EmployeeId": 1},
   	{
   		$set: {
   			"EmployeeName": "NewMartin2",
   			"EmployeeId": 2
   		}
   	}
   )
   ```

## 5. MongoDB 安全，监控和备份

### 5.1 MongoDB安全概述

MongoDB 能够定义数据库的安全性机制。默认情况下，我们不希望任何人都可以对MongoDB中的每个数据库进行开发访问操作，因此MongoDB中具有某种安全机制的要求很重要。

下面是在数据库中实现安全性的最佳操作：

1. 启用访问控制-创建用户，以便在访问MongoDB上的数据库时强制所有应用程序和用户具有某种身份的验证机制。
2. 配置基于角色的访问控制-优势可能需要对权限进行逻辑分组，这些分组可以包含在角色中，然后可以将用户分配给这些角色。
3. 尝试将MongoDB配置为使用某种加密协议，例如TLS或SSL。这些协议可用于加密在客户端和MongoDB环境之间流动的流量。
4. 配置审计-管理员通常需要知道谁在做什么，这有助于以后分析问题。最好的方法时在MongoDB中启用审计。
5. 使用单独的用户ID运行MongoDB服务器示例，该用户ID可以访问服务器环境上所需的资源。

### 5.2 MongoDB备份程序-Mongodump

使用MongoDB时，务必确保备份过程顺利，成功，以防止MongoDB中的数据由于任何原因而损坏。

1. 通过复制底层数据文件进行备份-这可能是简单的机制，所以要做的就是复制MongoDB所在的数据文件，理想情况是将其复制到另外一条服务器。
2. 使用mongodump备份数据库-mongodump工具从MongoDB数据库读取数据并创建高保真BSON文件。需要注意的是，如果数据量很大，Mongodump可能会占用大量资源，所以，为缓解这种情况，应该在服务服务器上运行此程序。
3. MongoDB Cloud Manager 备份-通过从MongoDB环境中读取操作日志数据来维持备份MongoDB副本集和分片群集。MongoDB Cloud Manager 可以通过存储操作日志数据来创建时间点恢复，以便随时为特定副本集或分片集群进行恢复。

### 5.3 MongoDB监控

监控是MongoDB最关键的管理活动之一。这是因为通过监控环境中可能出现的问题，以便更好地解决问题。

**下面是一些实施监控的例子：**

1. Mongostat 将告诉服务器上实际上发生了多少次数据库操作，例如插入，查询，更新，删除等。这将为服务器处理多少负载提供了一个参考，并将指示您是否需要服务器上的其他资源，或者可能需要其他服务器来分配负载。
2. mongotop跟踪并报告MongoDB实例的当前读写活动，并基于每个集合报告这些统计信息。
3. MongoDB提供了一个Web界面，可在一个简单的网页中显示诊断监视信息。
4. serverStatus命令或命令程序中的serverStatus() 返回数据库状态的概述，包括磁盘使用情况，内存使用情况，与MongoDB环境建立的连接等详细信息。

### 5.4 MongoDB索引和性能影响

1. 索引在任何数据库中都非常重要，可用于提高MongoDB中的搜索查询效率。如果您继续在文档中执行搜索，则最好在搜索条件中使用的字段上添加索引。
2. 尝试始终限制返回的查询结果的数量。假设在文档中有2个字段名称，只想从文档中看到2个字段。索引确保查询针对显示需要的2个字段，而不是全部字段。
3. 如果要查看某些字段值，则在查询中使用这些字段，如果不需要，请不要查询集合中的所有字段。

## 6. MongoDB创建用户并添加角色

### 6.1 创建管理员用户

通过使用createUser方法在MongoDB中创建用户管理员。

```sql
db.createUser(
	{
		user: "Guru99",
		pwd: "password",
		roles: [
			{role: "userAdminAnyDatabase", db: "admin"}
		]
	}
)
```

![image-20250318101857652](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318101857652.png)

**代码说明：**

1. 第一步指定需要创建的“username”和“password”
2. 第二步是为用户分配角色。由于它需要时数据库管理员，这种情况下，我们已分配给“userAdminAnyDatabase”角色。该角色允许用户对MongoDB中的所有数据库具有管理特权。
3. db参数指定admin数据库，它是MongoDB中一个特殊的元数据库，其中包含该用户的信息。

### 6.2 MongoDB为单个数据库创建用户

要创建将管理单个数据库的用户，我们可以使用上述相同的命令，但是我们只需要使用“userAdmin”选项

```sql
db.createUser(
	{
		user: "Employeeadmin",
		pwd: "password",
		roles:[
			{role: "userAdmin", db: "Employee"}
		]
	}
)
```

![image-20250318102518536](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318102518536.png)

**代码说明：**

1. 第一步是指定需要创建的“username”和“password”
2. 第二步是为用户分配一个角色，在这种情况下，由于需要成为数据库管理员，因此将其分配给“userAdmin”角色。该角色允许用户仅对db选项中指定的数据库具有管理特权。
3. db参数指定用户对应其具有管理特权的数据库

### 6.3 管理用户

首先了解我们需要定义的角色。MongoDB中提供了完整的角色列表。例如：有一个“read role”仅允许对数据库进行只读操作，然后有一个“readwrite”角色提供了对数据库符读写访问，这意味着用户可以发出插入，删除和更新该数据库中集合的命令。

```sql
db.createUser(
  {
	user: "Mohan",
  pwd: "password",
  roles: [
  	{
  		role: "read", db: "Marketing"
  	},
  	{
  		role: "readWrite", db: "Sales"
  	}
  ]
  }
)
```

![image-20250318103455648](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318103455648.png)

如上述代码，显示了创建一个名为Mohan的用户，并且在多个数据库中分配了多个角色。在上面的实例中，它被授予对“Marketing”数据库的只读权限，对“Sales"数据库具有readwrite权限。

## 7 使用Kerberos身份验证配置MongoDB：X.509证书

授权是为了确保客户端对系统的访问，而身份验证会检查客户端在被授权进入系统后在MongoDB中国的访问类型。

身份验证机制有很多种，以下是其中几种：

### 7.1 使用x.509证书的MongoDB身份验证

使用x.509证书对客户端进行身份验证-证书基本上是客户端和MongoDB服务器之间的信任签名。

因此，无需输入用户名和密码来连接服务器，而是在客户端和MongoDB服务器之间传递证书。客户端基本上将具有客户端证书，该证书将传递到服务器以对服务器进行身份验证。每个客户端证书对应一个MongoDB用户。因此，来自MongoDB的每个用户都必须拥有自己的证书才能通过MongoDB服务器进行身份验证。

为确保此功能有效，必须遵循以下步骤：

1. 必须从有效的第三方机构买有效的证书并将其安装在MongoDB服务器上。
2. 客户端证书必须具有以下属性（单个证书颁发机构（CA）必须同时为客户端和服务端颁发证书。客户端证书必须包含以下字段-keyUsage和extendedKeyUsage）
3. 连接到MongoDB服务器的每个用户都需要有一个单独的证书。

### 7.2 使用 Kerberos 的MongoDB身份验证

1. 在Windows上使用Kerberos身份验证配置MongoDB - Kerberos是大型客户端-服务器环境中使用的身份验证机制。这是一种非常安全的机制。其中密码只有在加密时允许使用。当然，MongoDB具有针对现有的基于Kerberos的系统进行身份验证的功能。

2. 启动mongoDB服务器进程

3. 启动MongoDB客户端进程并连接到MongoDB服务器

4. 在MongoDB中添加一个用户，该用户基本上是 $external 数据库的 Kerberos 主体名称。$external数据库是一个特殊的数据库，它告诉MongoDB根据Kerberos系统而不是其自己内部的系统对该用户进行身份验证。

   ```sql
   use $external
   db.createUser(
   	{
   		user: "user1@example.NET",
   		roles: [
   			{
   				role: "read", db: "Marketing"
   			}
   		]
   	}
   )
   ```

   ![image-20250318114612970](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318114612970.png)

5. 使用命令启动具有 Kerberos 支持的 mongoDB

   ```sh
   mongdbsh -auth -setParameter authenticationMechanisms=GSSAPI
   ```

   限制可以使用Kerberos用户和Kerberos身份验证连接到数据库。

## 8. MongoDB副本集

### 8.1 什么是 MongoDB 复制？

复制是指确保同一数据在多个 MongoDB DB Server上可用的过程。实现数据高可用性，这样做很有必要。

如果我们的主MongoDB服务器由于任何原因而关闭，将无法访问数据。但是，如果我们有规律地将数复制到另一台服务器，即使主服务器出现故障，也可以从另一台服务器访问数据。

复制到另一个目的是实现负载均衡。如果有许多用户连接到系统，而不是让所有用户都连接到一个系统。可以将用户连接到多台服务器，以使负载平均分配。

在MongoDB中，多个MongoDB服务器会被分组在称为副本集的集合中。副本集将具有一个主服务器，这些实例可主要用于所有读取操作。

### 8.2 副本集：使用 rs.initiate() 添加第一个成员

如上面所讲，要启用复制，我们首先需要查UN更加爱你MongoDB实例的副本集。

假设在我们的示例中，我们有3个服务器，分别称为ServerA，S erverB, ServerC，在此配置中，ServerA将成为我们的主服务器，ServerB和ServerC将成为我们的从服务器。下图是一个很好的展示。

![image-20250318140356840](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20250318140356840.png)

以下是创建副本集以及将第一个成员添加到副本集所需遵循的步骤。

1. 确保将要添加到副本集的所有 mongodb 实例都安装在不同的服务器上。这是为了确保即使一台服务器出现故障，其他服务也将可用，因此 mongodb 的其他实例也将可用。

2. 确保所有 mongodb 实例都可以相互连接。从 ServerA 发出以下2条命令

   ```sh
   mongo -host ServerB -port 27017
   mongo -host ServerC -port 27017
   ```

   同样，如果有更多从服务器，可以执行相同的操作。

3. 使用 replSet选项启动第一个 mongodb 实例。此选项为将成为此副本集一部分的所有服务器提供分组。

   ```sh
   mongo -replSet "Replicat1"
   ```

   其中 “Replicat1” 是副本集的名称。我们可以为副本集定义任何名称。

4. 现在，将第一台服务器添加到副本集，下一步是通过发出以下命令 rs.initiate() 来启动副本集。

5. 通过发出命令 rs.conf() 来验证副本集，以确保副本设置正确。

### 8.3 使用 rs.add() 添加辅助副本

只需使用 rs.add 命令就可以将从服务器添加到副本集。此命令输入从服务器名称，并将服务器添加到复制集中。

假设有 ServerA,ServerB,ServerC，并且 ServerA 被定义为副本集中的主服务器。要将ServerB和ServerC添加到副本集，执行命令。

rs.add("ServerB")   rs.add("SercerC")













