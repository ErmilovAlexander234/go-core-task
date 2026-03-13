package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 <- 4
	ch2 <- 5
	ch2 <- 6
	close(ch2)

	merged := MergeChannels(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 4, 5, 6}

	for _, val := range expected {
		found := false
		for _, r := range result {
			if r == val {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("value %v not found in merged channel", val)
		}
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d elements, got %d", len(expected), len(result))
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	close(ch1)
	close(ch2)

	merged := MergeChannels(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	if len(result) != 0 {
		t.Error("expected empty merged channel")
	}
}
