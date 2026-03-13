package main

import (
	"fmt"
	"math/rand"
	"time"
)

func originalSlice() []int {
	rand.Seed(time.Now().UnixNano())

	s := make([]int, 10)

	for i := range s {
		s[i] = rand.Intn(1000)
	}

	return s
}

func sliceExample(s []int) []int {
	var result []int

	for _, v := range s {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
}

func addElements(s []int, n int) []int {
	result := make([]int, len(s))
	copy(result, s)

	result = append(result, n)

	return result
}

func copySlice(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)
	return result
}

func removeElement(s []int, index int) []int {

	if index < 0 || index >= len(s) {
		return s
	}

	result := make([]int, 0, len(s)-1)

	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)

	return result
}

func main() {

	orig := originalSlice()
	fmt.Println("Original slice:", orig)

	even := sliceExample(orig)
	fmt.Println("Even numbers:", even)

	added := addElements(orig, 999)
	fmt.Println("After adding element:", added)

	copied := copySlice(orig)
	orig[0] = 1000

	fmt.Println("Modified original:", orig)
	fmt.Println("Copied slice:", copied)

	removed := removeElement(orig, 1)
	fmt.Println("After removing index 1:", removed)
}
