package main

import "fmt"

// maxSumFixedWindow - max sum of any contiguous subarray of length k.
// Returns 0 if k is invalid or nums is shorter than k.
func maxSumFixedWindow(nums []int, k int) int {
	if k <= 0 || len(nums) < k {
		return 0
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	best := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > best {
			best = sum
		}
	}
	return best
}

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	fmt.Printf("nums=%v k=%d maxSum=%d\n", nums, k, maxSumFixedWindow(nums, k))
}
