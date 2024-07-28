# 设计模式

## 1.Go 语言设计模式

基于golang的设计模式思想

参照：[https://github.com/senghoo/golang-design-pattern](https://github.com/senghoo/golang-design-pattern)

​	在软件领域，GOF（一般用来指经典名著《设计模式》的四位作者）首次系统化提出三大类共25种可复用的经典设计方案来解决常见的软件问题，为可复用软件奠定了一定的理论基础。

​	从总体上来说，这些设计模式分为创建型模式，行为型模式，结构型模式三大类。

![image-20240728104100698](https://picpoahu.oss-cn-chengdu.aliyuncs.com/images/image-20240728104100698.png)

## 2.创建型模式

创建型模式（Creational Pattern）提供了一种在创建对象的同时隐藏创建逻辑的方式，而不是使用 new 运算符直接实例化对象。

在这种类型的设计模式里，单例模式和工厂模式（具体包括简单工厂模式，抽象工厂模式和工厂方法模式3种）在Go项目比较常见。

### 2.1单例模式

​	单例模式（singleton pattern）是最简单的一个模式。在 Go 中，单例模式指的是全局只有一个实例，并且它负责创建自己的对象。单例模式有减少内存和系统资源开销，防止多个实例产生冲突等优点。	

​	因为单例模式保证了实例全局的唯一性，而且只被初始化一次，所以比较适合全局共享一个实例，且只需要被实例化一次的场景，例如数据库实例，全局配置，全局任务池等。

​	单例模式分为饿汉方式和懒汉方式。饿汉方式是指全局的单例实例在包加载时创建，而懒汉方式指全局的单例实例在第一次被使用时创建。

**饿汉方式**

```go
package singleton

// Hungry 单例模式 饿汉方式
type Hungry struct {
}

var hungry *Hungry = &Hungry{}

func GetInsOr() *Hungry {
    return hungry
}
```

**懒汉方式**

```go
package singleton

import "sync"

/*
1.懒汉方式是开源项目使用最多的方式，但它的缺点就是“非并发安全”，在实际使用中需要加锁。
2.为了解决懒汉方式非并发安全的问题，需要对实例进行加锁
3.使用 once.Do 可以确保 lazy 实例全局只被创建一次，once.Do 函数可以确保当同时创还能多个对象时，只被一个动作执行
*/
type Lazy struct{}

var lazy *Lazy
var mu sync.Mutex
var once sync.Once

// GetLazy 非并发安全实现
func GetLazy() *Lazy {
    if lazy == nil {
       lazy = &Lazy{}
    }
    return lazy
}

// GetLazySafe 加锁实现
func GetLazySafe() *Lazy {
    if lazy == nil {
       mu.Lock()
       if lazy == nil {
          lazy = &Lazy{}
       }
       mu.Lock()
    }
    return lazy
}

// GetLazyOnce 更加优雅的一种实现
func GetLazyOnce() *Lazy {
    once.Do(func() {
       lazy = &Lazy{}
    })
    return lazy
}
```

### 2.2工厂模式

**1.简单工厂模式：最常用，最简单。**

和 p := &Person{} 这种创建出来的实例相比，简单工厂模式可以确保我们常见实例所需要的参数，进而可以保证实例的方法可以按预期执行。

```go
package factory

import "fmt"

// Person 简单工厂模式
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet(){
    fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
    return &Person{
       Name: name,
       Age: age,
    }
}
```

**2.抽象工厂模式：返回的是接口，而不是结构体。**

```go
package factory

import "fmt"

type Person2 interface {
    Greet2()
}

type person2 struct {
    name string
    age  int
}

func (p person2) Greet2() {
    fmt.Printf("Hi! My name is %s", p.name)
}

// NewPerson2 返回一个接口，而非结构体
func NewPerson2(name string, age int) Person2 {
    return person2{
       name: name,
       age:  age,
    }
}
```

**3.工厂方法模式**：将对象的创建从由一个对县负责所有具体类的实例化，变成由一群子类负责对具体类的实例化，从而将过程解耦。

```go
package factory

import "fmt"

// Person3 工厂方法模式
type Person3 struct {
    name string
    age  int
}

func NewPersonFactory(age int) func(name string) Person3 {
    return func(name string) Person3 {
       return Person3{
          name: name,
          age:  age,
       }
    }
}

func main() {
    //创建具有默认年龄的工厂
    newBaby := NewPersonFactory(1)
    baby := newBaby("john")
    fmt.Printf("john is %v", baby)

    newTeenager := NewPersonFactory(16)
    teen := newTeenager("jill")
    fmt.Printf("teen is %v", teen)
}
```

## 3.行为型模式

### 3.1策略模式

策略模式定义了一组算法，将每个算法都封装起来。并使它们之间可以互换。

```go
package strategy

// 策略模式

// IStrategy 定义一个策略类
type IStrategy interface {
    do(int, int) int
}

//策略实现：加
type add struct {}

func (*add) do(a, b int) int {
    return a + b
}

//策略实现：减
type reduce struct {}

func (*reduce) do(a, b int) int{
    return a - b
}

// Operator 具体的策略执行者
type Operator struct {
    strategy IStrategy
}

//设置策略
func (operator *Operator) setStrategy(strategy IStrategy){
    operator.strategy = strategy
}

//调动策略的方法
func (operator *Operator) calculate (a, b int) int{
    return operator.strategy.do(a, b)
}

// 随意更改策略，而不会影响 Operator 的所有实现
func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add：", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(3, 1)
	fmt.Println("add：", result)
}
```

### 3.2模版模式

​	模版模式（Template）定义了一个操作中算法的骨架，并将一些步骤延迟到子类中。这些方法可以让子类在不改变一个算法结构的情况下，重新定义该算法的某些特定步骤。

​	简单来说，模版模式就是将一个类中能够公共使用的方法就放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为模版，让开发者去填充需要填充的地方。

```go
package template

import "fmt"

type Cooker interface {
    fire()
    cooke()
    outFire()
}

// CookMenu 类似于一个抽象类
type CookMenu struct {}

func (CookMenu) fire() {
    fmt.Println("开火")
}

//做菜交给具体的子类实现
func (CookMenu) cooke() {
    
}

func (CookMenu) outFile() {
    fmt.Println("关火")
}

//封装具体步骤
func doCook(cook Cooker)  {
    cook.fire()
    cook.cooke()
    cook.outFire()
}

type XiHongShi struct {
    CookMenu
}

func (XiHongShi) cooke() {
    fmt.Println("做西红柿")
}

type ChaoJiDan struct {
    CookMenu
}

func (ChaoJiDan) cooke() {
    fmt.Println("做炒鸡蛋")
}
```

## 4.结构型模式

结构型模式的特点是关注对象之间的通信。

### 4.1代理模式

代理模式可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问。以下代码是一个代理模式的实现。

```go
package proxy

import "fmt"

// Seller 代理模式
type Seller interface {
    sell(name string)
}

// Station 火车站
type Station struct {
    stock int // 库存
}

func (station *Station) sell(name string) {
    if station.stock > 0 {
       station.stock--
       fmt.Printf("代理点中：%s买了一张票，剩余：%d \n", name, station.stock)
    } else {
       fmt.Println("票已售空")
    }
}

// StationProxy 火车代理点
type StationProxy struct {
    station *Station //持有一个火车站对象
}

// 下面代码中，StationProxy代理了 Station，代理类中持有被代理类对象，并且和代理类对象实现了同一接口
func (proxy *StationProxy) sell(name string) {
    if proxy.station.stock > 0 {
       proxy.station.stock--
       fmt.Printf("代理点中：%s买了一张票，剩余：%d \n", name, proxy.station.stock)
    } else {
       fmt.Println("票已售空")
    }
}
```

### 4.2选项模式

```golang
package option

import "time"

type Connection struct {
    addr    string
    cache   bool
    timeout time.Duration
}

const (
    defaultTimeout = 10
    defaultCaching = false
)

type options struct {
    timeout time.Duration
    caching bool
}

type Option interface {
    apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
    f(o)
}

func WithTimeout(t time.Duration) Option {
    return optionFunc(func(o *options) {
       o.timeout = t
    })
}

func WithCaching(cache bool) Option {
    return optionFunc(func(o *options) {
       o.caching = cache
    })
}

// NewConnect 创建一个新的Connection
func NewConnect(addr string, opts ...Option) (*Connection, error) {
    options := options{
       timeout: defaultTimeout,
       caching: defaultCaching,
    }

    for _, o := range opts {
       o.apply(&options)
    }

    return &Connection{
       addr:    addr,
       cache:   options.caching,
       timeout: options.timeout,
    }, nil
}
```













