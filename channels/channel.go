package channel

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; i < 10 ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func pong(c chan string) {
	for i := 1; i < 10 ; i++ {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func print(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func Run() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go print(c)

	time.Sleep(10 * time.Second)
}