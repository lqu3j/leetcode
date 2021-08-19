package main

import "log"

func main() {
	log.Println(plusOne([]int{0}))
}

func plusOne(digits []int) []int {
	var (
		quotient = 1
		length   = len(digits)
	)
	for i := length - 1; i >= 0; i-- {
		tmp := (digits[i] + quotient) / 10
		if tmp > 0 {
			digits[i] = (digits[i] + quotient) % 10
			quotient = tmp
		} else {
			digits[i] = digits[i] + quotient
			quotient = 0
			break
		}
	}
	log.Println(quotient, digits)
	if quotient > 0 {
		return append([]int{quotient}, digits...)
	}
	return digits
}
