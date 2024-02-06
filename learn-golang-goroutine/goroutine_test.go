package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutineMul(t *testing.T) {

	channel := make(chan int)
	go MultiplyNum(10, 10, channel)
	fmt.Println("Running")

	result := <-channel
	fmt.Println(result)

	time.Sleep(1 * time.Second)
}

func TestGoroutineLoop(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
