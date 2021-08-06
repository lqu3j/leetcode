package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println(multiply("9133", "0"))
}

func multiply(num1 string, num2 string) string {
	var (
		sum     string
		num1Len = len(num1) - 1
		num2Len = len(num2) - 1
	)
	if num1Len == 0 && num1[0] == '0' {
		return "0"
	}

	if num2Len == 0 && num2[0] == '0' {

		return "0"
	}

	for i := num1Len; i >= 0; i-- {
		for j := num2Len; j >= 0; j-- {
			tmp := strconv.Itoa(int((num1[i]-'0')*(num2[j]-'0'))) + strings.Repeat("0", num1Len-i+num2Len-j)
			sum = sumStr(sum, tmp)
		}
	}

	return sum
}

func sumStr(num1 string, num2 string) string {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	index1 := len(num1) - 1
	index2 := len(num2) - 1
	carray := 0
	for {
		if index1 >= 0 && index2 >= 0 {
			sum := int(num1[index1]-'0'+num2[index2]-'0') + carray
			quotient := sum % 10
			carray = sum / 10
			buf.Write([]byte{'0' + byte(quotient)})
		} else if index1 < 0 && index2 >= 0 {
			sum := int(num2[index2]-'0') + carray
			quotient := sum % 10
			carray = sum / 10
			buf.Write([]byte{'0' + byte(quotient)})
		} else if index1 >= 0 && index2 < 0 {
			sum := int(num1[index1]-'0') + carray
			quotient := sum % 10
			carray = sum / 10
			buf.Write([]byte{'0' + byte(quotient)})
		} else {
			break
		}
		index1--
		index2--
	}
	if carray != 0 {
		buf.Write([]byte{'0' + byte(carray)})
	}
	return reverse(buf.String())
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
