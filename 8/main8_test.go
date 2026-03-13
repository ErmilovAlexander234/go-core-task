package main

import (
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {

	wg := NewWaitGroup()

	done := false

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		done = true
	}()

	wg.Wait()

	if !done {
		t.Error("goroutine did not finish")
	}
}

func TestMultipleGoroutines(t *testing.T) {

	wg := NewWaitGroup()

	count := 0
	n := 5

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()

	if count != n {
		t.Errorf("expected %d got %d", n, count)
	}
}
