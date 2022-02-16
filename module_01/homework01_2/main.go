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
			}
		}
	}()
	//生产者
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		var i1 int
		for range ticker.C {
			select {
			case <-done:
				fmt.Printf("%v: 生产者 child process interrupt...\n", time.Now())
				return
			default:
				for i := 0; i <= 10; i++ {
					i1 = int(time.Now().Second()*100 + i)
					message <- i1
					fmt.Printf("%v: len message=%d; send message: %v\n", time.Now(), len(message), i1)
				}
			}
		}
	}()
	time.Sleep(10 * time.Second)
	close(done)
	fmt.Println("main process exit!")
	time.Sleep(2 * time.Second)
}
