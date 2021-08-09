package main

import "log"

func main() {
	log.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	results := [][]int{}
	if target <= 0 {
		return results
	}

	for i := 0; i < len(candidates); i++ {
		sub := target - candidates[i]
		if sub < 0 {
			return [][]int{}
		} else if sub == 0 {
			return [][]int{{candidates[i]}}
		} else {
			tmp := combinationSum(candidates[i:], target-candidates[i])
		}

		// tmp := combinationSum(candidates[i:], target-candidates[i])

		// if len(tmp) == 0 {
		// 	tmp = append(tmp, []int{candidates[i]})
		// } else {
		// 	for j := range tmp {
		// 		tmp[j] = append(tmp[j], candidates[i])
		// 	}
		// }

		// if target-candidates[i] == 0 {
		// 	results = append(results, tmp...)
		// }
	}
	return results
}
