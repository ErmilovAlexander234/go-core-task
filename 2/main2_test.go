package main

import "testing"

func TestOriginalSlice(t *testing.T) {
	s := originalSlice()

	if len(s) != 10 {
		t.Error("slice length should be 10")
	}
}

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}

	result := sliceExample(input)

	expected := []int{2, 4, 6}

	if len(result) != len(expected) {
		t.Error("wrong even slice length")
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Error("even numbers filtering failed")
		}
	}
}

func TestAddElements(t *testing.T) {
	s := []int{1, 2, 3}

	result := addElements(s, 10)

	if result[len(result)-1] != 10 {
		t.Error("element was not added")
	}

	if len(result) != 4 {
		t.Error("slice size incorrect after add")
	}
}

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3}

	copy := copySlice(s)

	s[0] = 100

	if copy[0] == 100 {
		t.Error("copy should not change when original changes")
	}
}

func TestRemoveElement(t *testing.T) {
	s := []int{10, 20, 30, 40}

	result := removeElement(s, 1)

	expected := []int{10, 30, 40}

	if len(result) != len(expected) {
		t.Error("wrong slice length after remove")
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Error("element was not removed correctly")
		}
	}
}
