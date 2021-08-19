package main

import "strings"

func main() {

}

func lengthOfLastWord(s string) int {
	array := strings.Split(s, " ")
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == "" {
			continue
		} else {
			return len(array[i])
		}
	}
	return 0
}
