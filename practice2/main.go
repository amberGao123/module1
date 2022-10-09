package main

import (
	"time"
)

func main() {

	queue := make(chan int, 10)

	go func(a chan int) {
		for {
			time.Sleep(time.Second)
			a <- 1
		}
	}(queue)

	go func(b chan int) {
		for {
			time.Sleep(time.Second)
			<-b

		}
	}(queue)

	select {}
}
