package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}

func findMaxForm(strs []string, m int, n int) int {
	cache := make([][]int, m+1, m+1)
	for i := range cache {
		cache[i] = make([]int, n+1, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zero := strings.Count(strs[i], "0")
		one := strings.Count(strs[i], "1")

		if zero > m || one > n {
			continue
		}
		cache[zero][one]++
	}
}
