package main

import (
	"fmt"
	"sync"
)

//3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，
//按照cat、dog、fish顺序打印在屏幕上100次
/*type printAnimal interface {
	print() string
}
*/
type cat struct {
	name string
}

type dog struct {
	name string
}

type fish struct {
	name string
}

func printCat(c *cat, ch1 chan struct{}, ch2 chan struct{}, sw *sync.WaitGroup) {
	defer sw.Done()
	for i := 0; i < 100; i++ {
		<-ch1
		fmt.Printf("%s---%d\n", c.name, i)
		ch2 <- struct{}{}
	}
}

func printDog(d *dog, ch2 chan struct{}, ch3 chan struct{}, sw *sync.WaitGroup) {
	defer sw.Done()
	for i := 0; i < 100; i++ {
		<-ch2
		fmt.Printf("%s---%d\n", d.name, i)
		ch3 <- struct{}{}
	}
}

func printFish(f *fish, ch3 chan struct{}, ch1 chan struct{}, sw *sync.WaitGroup) {
	defer sw.Done()
	for i := 0; i < 100; i++ {
		<-ch3
		fmt.Printf("%s---%d\n", f.name, i)
		ch1 <- struct{}{}
	}
}
func main() {

	n1 := make(chan struct{}, 1)
	n2 := make(chan struct{}, 1)
	n3 := make(chan struct{}, 1)
	sw := &sync.WaitGroup{}
	sw.Add(3)
	n1 <- struct{}{}

	cat := &cat{
		name: "cat",
	}

	dog := &dog{
		name: "dog",
	}

	fish := &fish{
		name: "fish",
	}

	//time.Sleep(time.Second * 5)
	go printCat(cat, n1, n2, sw)
	go printDog(dog, n2, n3, sw)
	go printFish(fish, n3, n1, sw)

	sw.Wait()
	fmt.Println("函数打印结束")
}
