package main

import (
	"fmt"
	"sync"
)

func MergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	copyChannel := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			out <- v
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go copyChannel(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 7; i <= 9; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merged := MergeChannels(ch1, ch2, ch3)

	for v := range merged {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
