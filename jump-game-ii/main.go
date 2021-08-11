package main

import "log"

func main() {
	log.Println(jump([]int{1, 2, 1, 1, 1}))
}

func jump(nums []int) int {
	length := len(nums)
	index := 0
	count := 0
	end := 0

	for index < length-1 {
		count++
		end = length - 1
		if index+nums[index] >= end {
			break
		} else {
			end = index + nums[index] + 1
		}
		index = getmax(nums, index+1, end)
	}

	return count
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
