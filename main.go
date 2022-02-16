package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func addZeros(a string, time int) string {
	for ; time > 0; time-- {
		a = "0" + a
	}
	return a
}

func Normalize(a, b string) (string, string) {
	if len(a) > len(b) {
		return a, addZeros(b, len(a)-len(b))
	}
	return addZeros(a, len(b)-len(a)), b
}

func LongAdd(a, b string) string {
	var result string
	var num uint8
	a, b = Normalize(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		num = (a[i] - '0') + (b[i] - '0') + num/10
		result = string((num%10)+'0') + result
	}
	if num/10 != 0 {
		result = "1" + result
	}
	return result
}

func Byte2Number(a byte) int {
	return strings.Index(base36, string(a))
}

func LongAddWithBase(a, b string, base int) string {
	var result string
	var add int
	a, b = Normalize(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		num := Byte2Number(a[i]) + Byte2Number(b[i]) + add
		add = num / base
		result = string(base36[num%base]) + result
	}
	if add != 0 {
		result = "1" + result
	}
	return result
}
func LongFibonacci(n int) string {
	prev, now := "0", "1"
	for ; n > 0; n-- {
		prev, now = now, LongAdd(now, prev)
	}
	return prev
}

func main() {
	fmt.Println(LongAdd("123", "2"))
}
