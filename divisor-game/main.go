package main

import "log"

func main() {
	log.Println(divisorGame(6))
}

func divisorGame(n int) bool {
	return n%2 == 0
}
