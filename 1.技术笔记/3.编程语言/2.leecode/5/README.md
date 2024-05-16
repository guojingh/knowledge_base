## 题目：3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，按照cat、dog、fish顺序打印在屏幕上100次
注意：
- 3个函数
- 3个goroutine
- 交替打印，顺序：cat dog fish
- 打印次数：100次

实现：
- 本题目使用了 channel + sync.waitGroup 的方式进行实现
- channel 负责发送消息，供三个 goroutine 进行通信，sync.waitGroup 负责等待3个goroutine的完成
- 作者水平有限，实现较为简单