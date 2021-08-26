package main

import (
	"fmt"
	"log"
)

func main() {

	array := []byte("1245")
	array = insert(array, 2, '3')
	log.Println(string(array))

	// log.Println(string(copyArray(array, len(array), 4)))

	// log.Println(getPermutation(1, 1))
}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	array := []byte("123456789")
	array = array[:n]
	return calc("", array, &k)
}

func calc(prefix string, array []byte, k *int) string {
	if len(array) == 2 {
		*k--
		if *k == 0 {
			return fmt.Sprintf("%s%s", prefix, string(array))
		}

		array[0], array[1] = array[1], array[0]

		*k--
		if *k == 0 {
			return fmt.Sprintf("%s%s", prefix, string(array))
		}
		array[0], array[1] = array[1], array[0]
		return ""
	}

	for i := 0; i < len(array); i++ {
		value := array[i]
		array = remove(array, i)
		results := calc(fmt.Sprintf("%s%c", prefix, value), array, k)
		if results != "" {
			return results
		}
		array = insert(array, i, value)
	}

	return ""
}

func remove(array []byte, index int) []byte {
	if index+1 >= len(array) {
		return array[:index]
	}
	return append(array[:index], array[index+1:]...)
}

func insert(array []byte, index int, value byte) []byte {
	return append(array[:index], append([]byte{value}, array[index:]...)...)
}
