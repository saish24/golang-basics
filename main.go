package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 30)
	go sendInTwoSeconds(ch)
	go receiveInThreeSeconds(ch)
	time.Sleep(20 * time.Second)
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func sendInTwoSeconds(channel chan<- int) {
	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Println("Sending ", i)
		time.Sleep(1 * time.Second)
	}
}

func receiveInThreeSeconds(channel <-chan int) {
	for i := 0; i < 10; i++ {
		j := <-channel
		fmt.Println("Received ", j)
		time.Sleep(3 * time.Second)
	}
}
