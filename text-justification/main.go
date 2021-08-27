package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
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
				results = append(results, join(words[begin:index], sum, maxWidth))
				begin = index
				sum = wordLen
			}
		}
		index++
	}
	if begin < length {
		results = append(results, strings.Join(words[begin:length], " ")+strings.Repeat(" ", maxWidth-sum))
	}
	return results
}

func join(array []string, currentLen int, maxWidth int) string {
	length := len(array)
	subLength := maxWidth - currentLen

	if length == 1 {
		return fmt.Sprintf("%v%v", array[0], strings.Repeat(" ", subLength))
	}
	if subLength == 0 {
		return strings.Join(array, " ")
	}

	quotient := subLength / (length - 1)
	remainder := subLength % (length - 1)

	buf := bytes.NewBuffer(make([]byte, 0, maxWidth))
	for i := 0; i < length; i++ {
		if i < length-1 {
			count := quotient + 1
			if remainder > 0 {
				count++
			}
			remainder--
			fmt.Fprintf(buf, "%v%v", array[i], strings.Repeat(" ", count))
		} else {
			fmt.Fprintf(buf, "%v", array[i])
		}
	}

	return buf.String()
}
