package main

func main() {
	insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
}

func insert(intervals [][]int, newInterval []int) [][]int {
	index := -1
	for i := 0; i < len(intervals); i++ {
		if newInterval[0] <= intervals[i][0] {
			index = i
			break
		}
	}
	if index == -1 {
		intervals = append(intervals, newInterval)
	} else {
		tmp := append([][]int{}, intervals[index:]...)
		intervals = append(intervals[:index], newInterval)
		intervals = append(intervals, tmp...)
	}

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
