package main

import "math"
import "bytes"
import "fmt"

func main() {
	fmt.Println(addBinary("11111", "11"))

}

func addBinary(a string, b string) string {
	aIndex, bIndex := len(a)-1, len(b)-1
	buf := bytes.NewBuffer(make([]byte, 0, int(math.Max(float64(aIndex), float64(bIndex)))+1))
	carray := 0

	for aIndex >= 0 && bIndex >= 0 {
		sum := int(a[aIndex]-'0') + int(b[bIndex]-'0') + carray
		carray = sum / 2
		fmt.Fprintf(buf, "%c", '0'+sum%2)
		aIndex--
		bIndex--
	}

	for aIndex >= 0 {
		sum := int(a[aIndex]-'0') + carray
		carray = sum / 2
		fmt.Fprintf(buf, "%c", '0'+sum%2)
		aIndex--
	}

	for bIndex >= 0 {
		sum := int(b[bIndex]-'0') + carray
		carray = sum / 2
		fmt.Fprintf(buf, "%c", '0'+sum%2)
		bIndex--
	}
	if carray != 0 {
		fmt.Fprintf(buf, "%c", '0'+carray)
	}
	return reverse(buf.Bytes())
}

func reverse(array []byte) string {
	begin := 0
	end := len(array) - 1

	for begin < end {
		array[begin], array[end] = array[end], array[begin]
		begin++
		end--
	}
	return string(array)
}
