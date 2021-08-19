package main

import (
	"log"
)

func main() {
	// a := []int{1, 2, 0}
	// a := []int{3, 4, -1, 1}
	// a := []int{7, 8, 9, 11, 12}
	// a := []int{2, 1}
	// a := []int{-1, -2, -60, 40, 43}
	// a := []int{1, 2, 6, 3, 5, 4}
	// a := []int{11, 1, 6, 11, 5, 5, -6, 9, -3, 9, 5, 4, 2, -8, 16, -6, -4, 2, 3}
	// a := []int{1, 1}
	// a := []int{41, 16, 65, 7, 56, 5, 64, 14, 67, 2, 3, 20, -10, 25, -8, 10, 2, 59, 9, 6, 43, 4, 22, 50, 41, 37, 5, 21, -2, 34, 34, 60, 22, 6, 4, 27, 1, 61, 45, 38, 31, 62, 63, -10, 4, 60, 21, 4, 4, 33, 40, 26, 16, 34, 51, 1, 5, 10, 20, -1, 54, 44, 3, 21, 63, 40, 4, 4}
	a := []int{-2, 10, -6, -3, 15, 17, -8, -1, -1, 17, 4, 9, -2, -5, 6, -5, -1, -1, 19, -1, -1, 17, -9}
	log.Println(firstMissingPositive(a))
}

func firstMissingPositive(nums []int) int {
	index := 0
	length := len(nums)

	// 处理所有<=0 以及 > length的数据
	for i := range nums {
		if nums[i] <= 0 || nums[i] > length {
			nums[i] = 0
		}
	}

	for index < length {
		v := nums[index]
		if v <= 0 {
			index++
			continue
		}
		if nums[v-1] >= 0 {
			nums[v-1], nums[index] = nums[index], nums[v-1]
		} else {
			index++
		}
		nums[v-1] = -2
	}

	for i := range nums {
		if nums[i] != -2 {
			return i + 1
		}
	}

	return length + 1
}
