package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}
	minArray := make([]int, len(cost), len(cost))

	for i := len(cost) - 1; i >= 0; i-- {
		if i >= len(cost)-2 {
			minArray[i] = cost[i]
		} else {
			minArray[i] = cost[i] + min(minArray[i+1], minArray[i+2])
		}
	}
	return min(minArray[0], minArray[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
