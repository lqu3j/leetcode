package main

import "log"

func main() {
	log.Println(isSubsequence("", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) {

		for j < len(t) && t[j] != s[i] {
			j++
		}

		if j >= len(t) {
			break
		}

		i++
		j++
	}

	return i == len(s)
}
