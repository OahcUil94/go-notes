package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func (i int) {
			// time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			time.Sleep(time.Second * 20)
			c.Broadcast()
		}(i)
	}

	// 注意, Wait方法的内部实现是在把waiter加入到队列中之后, 进行了Unlock解锁, 所以在使用Wait的时候一定要加锁
	// 在调用了Broadcast之后, Wait就会结束, 所以如果要保证所有的协程都执行结束, 那就要进行一定条件判断的等待
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有运动员是否就绪
	log.Println("所有运动员都准备就绪, 比赛开始, 3, 2, 1......")
}