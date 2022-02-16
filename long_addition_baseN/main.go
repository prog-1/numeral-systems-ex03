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

func NumberExistsInBaseN(num string, base int) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}

func LongAdd(a, b string, base int) string {
	var s string
	var add uint8 // uint8 = byte
	a, b = MakeNormal(a, b)
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if 2 < base && base <= 36 && NumberExistsInBaseN(a, base) && NumberExistsInBaseN(b, base) {
		for i := len(a) - 1; i >= 0; i-- {
			num := (a[i] - '0') + (b[i] - '0') + add // '0' = 48
			add = num / uint8(base)
			s = string(base36[num%uint8(base)]) + s
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
	var base int
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	if LongAdd(a, b, base) != "-1" {
		fmt.Println("The sum of a and b in an arbitary numeral system:", LongAdd(a, b, base))
	} else {
		fmt.Println("ERROR: Incorrect base")
	}
}
