package main

import (
	"log"
	"sort"
)

func main() {
	array := []int{1, 0}
	merge(array, 1, []int{1}, 1)

	log.Println(array)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Slice(nums1[:m], func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	index1 := 0
	index2 := 0
	length := m
	count := 0
	for index1 < m+n && index2 < n {
		log.Println(nums1, index1, nums2[index2])

		if nums1[index1] >= nums2[index2] {
			if count < m {
				move(nums1, length, index1)
			}
			nums1[index1] = nums2[index2]
			length++

			index1++
			index2++
		} else {
			if count < m {
				count++
				index1++
			} else {
				nums1[index1] = nums2[index2]
				index1++
				index2++
			}
		}
	}
}

func move(array []int, length, index int) {
	for i := length - 1; i >= index; i-- {
		array[i+1] = array[i]
	}
}
