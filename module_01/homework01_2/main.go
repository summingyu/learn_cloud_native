package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	message := make(chan int, 10)
	done := make(chan bool)
	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)
	defer close(message)
	// 消费者
	for i := 0; i <= 20; i++ {
		switch r.Intn(2) {
		case 0:
			go func(nums int) {
				fmt.Printf("消费者No.%d 加入战斗\n", nums)
				ticker := time.NewTicker(1 * time.Second)
				for range ticker.C {
					select {
					case <-done:
						fmt.Printf("消费者No.%d: 消费者child process interrupt...\n", nums)
						return
					case i := <-message:
						fmt.Printf("消费者No.%d:len message=%d; receive message: %v\n", nums, len(message), i)
					default:
						fmt.Printf("消费者No.%d:len message=%d; message is empty\n", nums, len(message))
					}
				}
			}(i)
		case 1:
			//生产者
			go func(nums int) {
				fmt.Printf("生产者No.%d 加入战斗\n", nums)
				ticker := time.NewTicker(1 * time.Second)
				var i1 int
				for range ticker.C {
					i1 = int(time.Now().Minute() * 100 + time.Now().Second())
					select {
					case <-done:
						fmt.Printf("生产者No.%d: 生产者 child process interrupt...\n", nums)
						return
					case message <- i1:
						fmt.Printf("生产者No.%d: len message=%d; send message: %v\n", nums, len(message), i1)
					default:
						fmt.Printf("生产者No.%d: len message=%d; message is full\n", nums, len(message))
					}
				}
			}(i)
		}
		time.Sleep(1 * time.Second)
	}
	// 停顿一秒用于验证队列为空时消费者是否走default
	time.Sleep(20 * time.Second)
	close(done)
	fmt.Println("main process exit!")
	time.Sleep(5 * time.Second)
}
