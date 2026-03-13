package main

import "fmt"

func checkIntersection(a, b []int) (bool, []int) {
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}

	intersection := make(map[int]struct{})
	for _, v := range b {
		if _, exists := m[v]; exists {
			intersection[v] = struct{}{}
		}
	}

	result := make([]int, 0, len(intersection))
	for v := range intersection {
		result = append(result, v)
	}

	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 42}
	b := []int{64, 3, 45, 43}

	ok, inter := checkIntersection(a, b)
	fmt.Println(ok, inter)
}
