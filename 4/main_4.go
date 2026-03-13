package main

import "fmt"

func Difference(slice1, slice2 []string) []string {
	set := make(map[string]struct{})
	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	var result []string
	for _, v := range slice1 {
		if _, exists := set[v]; !exists {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	diff := Difference(slice1, slice2)
	fmt.Println(diff)
}
