package main

import "fmt"

func main() {
	fmt.Print(climbStairs(10))
}

func climbStairs(n int) int {
	cache := make([]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		if i <= 2 {
			cache[i] = i
		} else {
			cache[i] = cache[i-1] + cache[i-2]
		}
	}
	if len(cache) > 0 {
		return cache[len(cache)-1]
	} else {
		return 0
	}
}
