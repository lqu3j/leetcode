package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	results := [][]int{}
	calc(candidates, []int{}, target, &results)
	return results
}

func calc(candidates []int, array []int, target int, results *[][]int) {

	if target == 0 {
		tmp := make([]int, len(array), len(array))
		copy(tmp, array)
		*results = append(*results, tmp)
	}

	for i := 0; i < len(candidates); i++ {
		sub := target - candidates[i]

		if sub < 0 {
			continue
		} else {

			calc(candidates[i:], append(array, candidates[i]), sub, results)
		}
	}

}
