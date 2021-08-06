package main

import "log"

func main() {
	log.Println(searchRange([]int{}, 8))
}

func searchRange(nums []int, target int) []int {
	left := -1
	right := -1

	index := binarySearch(nums, 0, len(nums)-1, target)
	if index == -1 {
		return []int{left, right}
	}

	left = index
	right = index

	for left-1 >= 0 && nums[left-1] == target {
		left--
	}

	for right+1 < len(nums) && nums[right+1] == target {
		right++
	}

	return []int{left, right}

}

func binarySearch(nums []int, begin, end, target int) int {
	length := end - begin + 1
	if length == 0 {
		return -1
	}
	mid := begin + length/2

	if target == nums[mid] {
		return mid
	}

	if target > nums[mid] {
		return binarySearch(nums, mid+1, end, target)
	} else {
		return binarySearch(nums, begin, mid-1, target)
	}
}
