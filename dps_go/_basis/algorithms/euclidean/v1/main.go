package main

import "fmt"

// gcd - Euclidean algorithm: greatest common divisor of a and b.
// Works with non-negative ints; negatives are normalized via abs.
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	fmt.Printf("gcd(48, 18) = %d\n", gcd(48, 18))
	fmt.Printf("gcd(17, 13) = %d\n", gcd(17, 13))
}
