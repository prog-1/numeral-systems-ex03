package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}
func zeros(a string, time int) string {
	for i := time; i > 0; i-- {
		a = "0" + a
	}
	return a
}

func LongAdd(a, b string) (sum string) {
	var add byte
	if len(a) > len(b) {
		b = zeros(b, len(a)-len(b))
	} else {
		a = zeros(a, len(b)-len(a))
	}
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + add
		add = num / 10
		sum = string(num%10+'0') + sum
	}
	if add != 0 {
		sum = "1" + sum
	}
	return sum
}

func main() {
	base := 10
	var a, b string
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	if IsValidNumber(a, base) && IsValidNumber(b, base) {
		fmt.Printf("Sum of %v and %v = %v", a, b, LongAdd(a, b))
	} else {
		fmt.Printf("The number %v or number %v is represented incorrect in base%v", a, b, base)
	}
}
