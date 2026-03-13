package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1, slice2 []string
		want           []string
	}{
		{
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			slice1: []string{"a", "b", "c"},
			slice2: []string{"x", "y", "z"},
			want:   []string{"a", "b", "c"},
		},
		{
			slice1: []string{"a", "b"},
			slice2: []string{"a", "b"},
			want:   []string{},
		},
	}

	for _, tt := range tests {
		got := Difference(tt.slice1, tt.slice2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Difference(%v, %v) = %v; want %v", tt.slice1, tt.slice2, got, tt.want)
		}
	}
}
