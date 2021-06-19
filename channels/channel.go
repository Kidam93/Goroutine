package channel

import (
	"fmt"
	"time"
)


func ping(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
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
	print(c)
}