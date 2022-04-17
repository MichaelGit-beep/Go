package main

import (
	"fmt"
	"sync"
)

var wg1 = sync.WaitGroup{}

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sender(i, c)
		go receiver(c)
	}
	wg.Wait()
}

func sender(i int, c chan int) {
	c <- i
	wg.Done()
}

func receiver(c chan int) {
	i := <-c
	fmt.Printf("%v\n", i)
	wg.Done()
}
