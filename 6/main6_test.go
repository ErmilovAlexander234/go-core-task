package main

import (
	"testing"
)

func TestRandomGeneratorLength(t *testing.T) {
	ch := RandomGenerator(5)
	count := 0
	for range ch {
		count++
	}
	if count != 5 {
		t.Errorf("expected 5 values, got %d", count)
	}
}

func TestRandomGeneratorValues(t *testing.T) {
	ch := RandomGenerator(10)
	for v := range ch {
		if v < 0 || v >= 1000 {
			t.Errorf("value %d out of range", v)
		}
	}
}
