# 简单工厂模式

`golang`中一般没有构造函数这么一说，所以一般会定义 `new` 函数来初始化相关的类。`new` 函数返回接口时，就可以看做是一个简单工厂模式。

在golang中一般推荐的做法就是简单工厂模式。

`simple_test.go`中只调用了` NewAPI()` 的方法，而具体的实现细节封装在` simple.go`中，因此这就简单的体现了一种简单工厂模式的思想，在 `golang `中，这种设计模式也是最常见的。

注意：在执行` simple_test.go` 时，要使用下面命令：

```
go test -v .\simple.go .\simple_test.go/
```

这是因为，在 GOPATH 模式下，go test 会为指定的源码文件生成一个虚拟代码包 `"command-line-arguments"`，而 `simple_test.go` 调用了 `simple.go` 中的 `Say()` 函数并不属于代码包 `"command-line-arguments" `，编译不通过，错误自然就产生了。因此要对两个文件都进行指定，再进行测试。
