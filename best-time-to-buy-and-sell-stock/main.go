package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 7, 1, 7, 1, 8}))
}

func maxProfit(prices []int) int {
	var (
		min    = 0
		profit = 0
	)

	for i := 1; i < len(prices); i++ {
		sub := prices[i] - prices[min]
		if profit < sub {
			profit = sub
		}
		if prices[min] > prices[i] {
			min = i
		}
	}

	return profit
}
