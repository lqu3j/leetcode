package main

import (
	"log"
)

func main() {
	log.Println(search([]int{2, 3, 4, 5, 6, 7, 8, 9, 1}, 9))
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
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

	left := binarySearch(nums, begin, mid-1, target)
	if left == -1 {
		return binarySearch(nums, mid+1, end, target)
	} else {
		return left
	}

}
