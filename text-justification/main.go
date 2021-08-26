package main

import (
	"log"
)

func main() {
	array := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	for i := 0; i < len(array); i++ {
		log.Println(array[i], len(array[i]))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var (
		index, length = 0, len(words)
		results       = []string{}
	)
	begin := 0
	sum := 0
	for index < length {
		wordLen := len(words[index])
		if index == 0 {
			sum = wordLen
			begin = 0
		} else {
			if sum+wordLen+1 <= maxWidth {
				sum = sum + wordLen + 1
			} else {
				results = append(results, join(words[begin:index], maxWidth))
				sum = 0
				begin = index
			}
		}
		index++
	}
	if begin < length {
		results = append(results, join(words[begin:length], maxWidth))
	}
	return results
}

func join(array []string, maxWidth int) string {
	lengt := len(array)
}
