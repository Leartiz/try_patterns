package main

import "testing"

func Test_maxNonOverlappingActivities(t *testing.T) {
	tests := []struct {
		name  string
		input []Interval
		want  int
	}{
		{
			name:  "empty",
			input: nil,
			want:  0,
		},
		{
			name:  "single",
			input: []Interval{{1, 4}},
			want:  1,
		},
		{
			name: "from notes",
			input: []Interval{
				{1, 4}, {3, 5}, {0, 6}, {5, 7}, {8, 9}, {5, 9},
			},
			want: 3,
		},
		{
			name: "touching boundaries",
			input: []Interval{
				{1, 2}, {2, 3}, {3, 4},
			},
			want: 3,
		},
		{
			name: "all overlap",
			input: []Interval{
				{1, 5}, {2, 6}, {3, 7},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := append([]Interval(nil), tt.input...)
			got := maxNonOverlappingActivities(input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
