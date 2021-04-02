package main

import (
	"log"
)

func main() {
	log.Println(maxSubArray([]int{-2, 1}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if max < sum {
			max = sum
		}
	}
	return max
}
