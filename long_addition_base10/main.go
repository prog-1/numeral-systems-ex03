package main

import (
	"fmt"
	"strings"
)

func AddZeros(s string, count int) string {
	for ; count > 0; count-- {
		s = "0" + s
	}
	return s
}

func MakeNormal(a, b string) (string, string) {
	if len(a) > len(b) {
		b = AddZeros(b, len(a)-len(b))
	} else {
		a = AddZeros(a, len(b)-len(a))
	}
	return a, b
}

func NumberExistsInBase10(num string) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range num {
		if !strings.Contains(base36[:10], string(v)) {
			return false
		}
	}
	return true
}

func LongAdd(a, b string) string {
	var s string
	var add uint8 // uint8 = byte
	a, b = MakeNormal(a, b)

	if NumberExistsInBase10(a) && NumberExistsInBase10(b) {
		for i := len(a) - 1; i >= 0; i-- {
			num := (a[i] - '0') + (b[i] - '0') + add // '0' = 48
			add = num / 10
			s = string((num%10)+'0') + s // '0' = 48
		}
		if add != 0 {
			s = "1" + s
		}
	} else {
		s = "-1"
	}
	return s
}

func main() {
	var a, b string
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	if LongAdd(a, b) != "-1" {
		fmt.Println("The sum of a and b:", LongAdd(a, b))
	} else {
		fmt.Println("ERROR: At least one of the entered numbers does not exist in base 10")
	}
}
