package main

import "log"

func main() {
	for i := 0; i < 100; i++ {
		log.Println(i, mySqrt(i))
	}
}

func mySqrt(x int) int {
	var (
		begin, end = 0, x
		mid        = 0
		product    = 0
	)

	for begin <= end {
		mid = begin + (end-begin)/2
		product = mid * mid

		if product == x {
			return mid
		} else if product < x {
			begin = mid
			if end-begin == 1 {
				if end*end <= x {
					return end
				} else {
					return begin
				}
			}
		} else {
			end = mid
		}

	}

	return mid
}
