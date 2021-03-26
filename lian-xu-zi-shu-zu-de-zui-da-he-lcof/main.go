package main

import (
	"log"
	"math"
)

func main() {
	log.Println(maxSubArray([]int{-1, -1, 2, 3, 5, -10, 11, -1}))
}

func maxSubArray(nums []int) int {
	var (
		sum int
		max = math.MinInt32
	)

	for i := range nums {
		sum += nums[i]

		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
