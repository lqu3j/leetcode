package main

import "log"

func main() {
	log.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}
	index := 0
	end := 0

	for index < length-1 {
		end = length - 1
		if index+nums[index] >= end {
			return true
		} else {
			end = index + nums[index] + 1
		}
		if nums[index] > 0 {
			index = getmax(nums, index+1, end)
		} else {
			return false
		}

	}
	return false
}

func getmax(nums []int, begin, end int) int {
	maxIndex := begin

	for begin < end {
		if maxIndex+nums[maxIndex] <= begin+nums[begin] {
			maxIndex = begin
		}
		begin++
	}
	return maxIndex
}
