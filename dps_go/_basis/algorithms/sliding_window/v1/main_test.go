package main

import "testing"

func Test_maxSumFixedWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "classic",
			nums: []int{2, 1, 5, 1, 3, 2},
			k:    3,
			want: 9, // 5+1+3
		},
		{
			name: "k equals len",
			nums: []int{1, 2, 3},
			k:    3,
			want: 6,
		},
		{
			name: "k too large",
			nums: []int{1, 2},
			k:    3,
			want: 0,
		},
		{
			name: "empty",
			nums: nil,
			k:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSumFixedWindow(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
