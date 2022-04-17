package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		c1 <- "one"
		close(c1)
	}()
	go func() {
		c2 <- "two"
		close(c2)
	}()

	time.Sleep(100 * time.Millisecond)
	for {
		if i, ok := <-c1; ok {
			fmt.Println("received", i)
		} else if i, ok := <-c2; ok {
			fmt.Println("received", i)
		}
	}
}
