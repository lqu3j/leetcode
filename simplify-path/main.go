package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf(simplifyPath("/../abc"))
}

func simplifyPath(path string) string {
	array := strings.Split(path, "/")

	for i := 0; i < len(array); i++ {
		if array[i] == ".." {
			if i-1 >= 0 {
				array = append(array[:i-1], array[i+1:]...)
				i = i - 2
			} else {
				array = array[i+1:]
				i--
			}
		} else if array[i] == "." || array[i] == "" {
			array = append(array[:i], array[i+1:]...)
			i--
		}
	}
	return "/" + strings.Join(array, "/")
}
