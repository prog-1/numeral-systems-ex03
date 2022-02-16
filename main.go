package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func addZeros(a string, time int) string {
	for i := time; i > 0; i-- {
		a = "0" + a
	}
	return a
}

func Normalize(a, b string) (string, string) {
	if len(a) > len(b) {
		b = addZeros(b, len(a)-len(b))
	} else {
		a = addZeros(a, len(b)-len(a))
	}
	return a, b
}

func LongAdd(a, b string) string {
	var result string
	var add uint8
	a, b = Normalize(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + add
		if num > 9 {
			add = 1
		} else {
			add = 0
		}
		result = string((num%10)+'0') + result
	}
	if add != 0 {
		result = "1" + result
	}
	return result
}

func LongAddWithBase(a, b string, base int) string {
	var result string
	var add int
	a, b = Normalize(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		num := (strings.Index(base36, string(a[i]))) + (strings.Index(base36, string(b[i]))) + add
		if num >= base {
			add = 1
		} else {
			add = 0
		}
		result = string(base36[num%base]) + result
	}
	if add != 0 {
		result = "1" + result
	}
	return result
}
func LongFibonacci(n int) string {
	if n == 0 {
		return "0"
	}
	prev, now := "0", "1"
	for i := 1; i < n; i++ {
		prev, now = now, LongAdd(prev, now)
	}
	return now
}

func main() {
	fmt.Println(LongAddWithBase("123", "2", 5))
}
