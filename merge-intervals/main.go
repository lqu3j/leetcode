package main

import (
	"log"
	"sort"
)

func main() {
	//	array := [][]int{
	//		{1, 3},
	//		{2, 6},
	//		{8, 10},
	//		{15, 18},
	//	}
	//	array := [][]int{
	//		{1, 4},
	//		{4, 5},
	//	}
	array := [][]int{
		{2, 3},
		{1, 4},
	}
	log.Println(merge(array))
}

func merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	results := [][]int{}
	length := len(intervals)
	i := 0
	for i < length {
		next := i + 1

		for next < length {
			if intervals[i][1] >= intervals[next][0] {
				if intervals[i][1] <= intervals[next][1] {
					intervals[i][1] = intervals[next][1]
				}
			} else {
				break
			}
			next++
		}
		results = append(results, intervals[i])
		i = next
	}
	return results
}
