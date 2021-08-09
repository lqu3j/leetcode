package main

import "log"

func main() {
	log.Println(search("AAAAAABC", "AAAB"))
}

func search(s string, p string) int {
	var (
		next = next(p)
		i, j = 0, 0
	)
	for i < len(s) {
		if s[i] == p[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = next[j-1]
		}
		if j == len(p) {
			return i - len(p)
		}
	}
	return -1
}

func next(p string) []int {
	var (
		next       = []int{0}
		index, now = 1, 0
	)

	for index < len(p) {
		if p[now] == p[index] {
			now++
			index++
			next = append(next, now)
		} else if now != 0 {
			now = next[now-1]
		} else {
			next = append(next, 0)
			index++
		}
	}
	return next
}
