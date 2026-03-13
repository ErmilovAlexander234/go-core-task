package main

import "fmt"

func Pipeline(input <-chan uint8, output chan<- float64) {
	for val := range input {
		res := float64(val)
		output <- res * res * res
	}
	close(output)
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)
	go Pipeline(input, output)

	go func() {
		pipeBomb := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, v := range pipeBomb {
			input <- v
		}
		close(input)

	}()
	for val := range output {
		fmt.Println(val)
	}
}
