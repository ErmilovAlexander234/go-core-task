package main

import (
	"fmt"
	"time"
)

type WaitGroup struct {
	counter int
	sem     chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(n int) {
	wg.counter += n
}

func (wg *WaitGroup) Done() {
	wg.sem <- struct{}{}
}

func (wg *WaitGroup) Wait() {
	for i := 0; i < wg.counter; i++ {
		<-wg.sem
	}
}

func main() {

	wg := NewWaitGroup()

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			time.Sleep(time.Millisecond * 100)
			fmt.Println("goroutine", n)

		}(i)
	}

	wg.Wait()

	fmt.Println("All goroutines finished")
}
