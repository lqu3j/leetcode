package main

import (
	"log"
)

func main() {
	// array := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	array := []string{"+", "e", ".", "+e", "+.", "e+", "e.", ".+", ".e", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", "0e0", "46.e3"}
	for i := range array {
		log.Println(array[i], isNumber(array[i]))
	}
}

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	var (
		index    = 0
		length   = len(s)
		eIndex   = -1
		dotIndex = -1
	)

	for index < length {
		switch s[index] {
		case '+', '-':
			if index == 0 || s[index-1] == 'e' || s[index-1] == 'E' {
				index++
				if index >= length {
					return false
				}
			} else {
				return false
			}
		case 'e', 'E':
			if eIndex != -1 {
				return false
			}
			eIndex = index

			if index > 0 && ((s[index-1] >= '0' && s[index-1] <= '9') || s[index-1] == '.') {
				index++
				if index >= length {
					return false
				}
			} else {
				return false
			}
		case '.':
			if dotIndex != -1 || eIndex != -1 {
				return false
			}
			dotIndex = index
			if !(index-1 >= 0 && s[index-1] >= '0' && s[index-1] <= '9') && !(index+1 < length && s[index+1] >= '0' && s[index+1] <= '9') {
				return false
			}
			index++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			index++
		default:
			return false
		}
	}

	return true
}
