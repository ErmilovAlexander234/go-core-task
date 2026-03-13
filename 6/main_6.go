package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomGenerator(n int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch <- rand.Intn(1000)
		}
		close(ch)
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := RandomGenerator(10)

	for v := range ch {
		fmt.Println(v)
	}
}
