package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)
	defer close(message)
	// 消费者
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-done:
				fmt.Printf("%v: 消费者child process interrupt...\n", time.Now())
				return
			case i := <-message:
				fmt.Printf("%v:len message=%d; receive message: %v\n", time.Now(), len(message), i)
			default:
				fmt.Printf("%v:len message=%d; message is empty\n", time.Now(), len(message))
			}
		}
	}()
	// 停顿一秒用于验证队列为空时消费者是否走default
	time.Sleep(1 * time.Second)
	//生产者
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		var i1 int
		for range ticker.C {
			i1 = int(time.Now().Second()*100)
			select {
			case <-done:
				fmt.Printf("%v: 生产者 child process interrupt...\n", time.Now())
				return
			case message <- i1:
				fmt.Printf("%v: len message=%d; send message: %v\n", time.Now(), len(message), i1)
			default:
				fmt.Printf("%v: len message=%d; message is full\n", time.Now(), len(message))

			}
		}
	}()
	// 插入10个数到队列,模拟队列满时的情况
	for i := 0; i <= 10; i++ {
		message <-i
	}
	time.Sleep(20 * time.Second)
	close(done)
	fmt.Println("main process exit!")
	time.Sleep(2 * time.Second)
}
