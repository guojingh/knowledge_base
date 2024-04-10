package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

// 测试生成一个对象的单例模式
func TestSinglet(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

// 测试生成多个对象的单例模式
func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instance := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			//协程阻塞，等待 channel 被关闭才能继续运行
			<-start
			instance[index] = GetInstance()
			wg.Done()
		}(i)
	}

	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instance[i] != instance[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
