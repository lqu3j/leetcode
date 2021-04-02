package main

import "log"

func main() {
	log.Println(waysToStep(76))
}

func waysToStep(n int) int {
	cache := make([]int, n+1, n+1)

	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 || i == 2 {
			cache[i] = i
		} else if i == 3 {
			cache[i] = 4
		} else {
			cache[i] = (cache[i-1] + cache[i-2] + cache[i-3]) % 1000000007
		}
	}
	if len(cache) != 0 {
		return cache[len(cache)-1]
	}
	return 0
}
