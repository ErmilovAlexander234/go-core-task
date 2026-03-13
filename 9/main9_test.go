package main

import "testing"

func TestPipeline(t *testing.T) {

	input := make(chan uint8)
	output := make(chan float64)

	go Pipeline(input, output)

	go func() {
		testData := []uint8{1, 2, 3}
		for _, v := range testData {
			input <- v
		}
		close(input)
	}()

	expected := []float64{1, 8, 27}

	i := 0
	for v := range output {
		if v != expected[i] {
			t.Errorf("expected %v got %v", expected[i], v)
		}
		i++
	}
}
