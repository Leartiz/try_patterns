package main

import "testing"

func Test_gcd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "classic", a: 48, b: 18, want: 6},
		{name: "coprime", a: 17, b: 13, want: 1},
		{name: "equal", a: 7, b: 7, want: 7},
		{name: "zero", a: 15, b: 0, want: 15},
		{name: "order", a: 18, b: 48, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcd(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("gcd(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
