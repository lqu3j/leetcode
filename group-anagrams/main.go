package main

import (
	"sort"
)

func main() {
	strs := []string{""}

	groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	mapping := make(map[string][]string)
	results := [][]string{}
	for _, v := range strs {
		tmp := []byte(v)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		mapping[string(tmp)] = append(mapping[string(tmp)], v)
	}

	for _, v := range mapping {
		results = append(results, v)
	}
	return results
}
