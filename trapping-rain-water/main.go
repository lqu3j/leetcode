package main

import (
	"log"
)

func main() {
	log.Println(trap([]int{0, 7, 1, 4, 6, 1, 5, 1, 4, 1, 3}))
}

func trap(height []int) int {
	var (
		total       int
		array       []int
		left, right = 0, 1
	)

	for right < len(height) {
		if height[left] == 0 {
			left = right
		} else {
			if height[left] > height[right] {
				array = append(array, height[left]-height[right])
			} else {
				left = right
				total += sum(array)
				array = []int{}
			}
		}
		right++
	}
	return total + remain(array)
}

func sum(array []int) int {
	sum := 0

	for i := range array {
		sum += array[i]
	}
	return sum
}

func remain(array []int) int {
	remain := 0
	last := len(array) - 1

	for i := last - 1; i >= 0; i-- {
		sub := array[i] - array[last]
		if sub > 0 {
			remain += sub
		} else {
			last = i
		}
	}
	return remain
}
