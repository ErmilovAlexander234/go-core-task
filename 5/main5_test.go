package main

import (
	"reflect"
	"testing"
)

func TestCheckIntersection(t *testing.T) {
	tests := []struct {
		a, b     []int
		wantBool bool
		want     []int
	}{
		{
			a:        []int{65, 3, 58, 678, 42},
			b:        []int{64, 3, 45, 43},
			wantBool: true,
			want:     []int{3},
		},
		{
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			wantBool: false,
			want:     []int{},
		},
		{
			a:        []int{1, 2, 2, 3},
			b:        []int{2, 3, 3},
			wantBool: true,
			want:     []int{2, 3},
		},
	}

	for _, tt := range tests {
		gotBool, got := checkIntersection(tt.a, tt.b)
		if gotBool != tt.wantBool {
			t.Errorf("expected bool %v, got %v", tt.wantBool, gotBool)
		}

		// проверка содержимого пересечения (порядок не важен)
		gotMap := make(map[int]struct{}, len(got))
		for _, v := range got {
			gotMap[v] = struct{}{}
		}
		wantMap := make(map[int]struct{}, len(tt.want))
		for _, v := range tt.want {
			wantMap[v] = struct{}{}
		}

		if !reflect.DeepEqual(gotMap, wantMap) {
			t.Errorf("expected intersection %v, got %v", tt.want, got)
		}
	}
}
