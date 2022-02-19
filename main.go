package main

import (
	"fmt"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LongAdd(a, b string) (s string) {
	if len(a) < len(b) {
		a, b = b, a
	}
	iB := len(b)
	var r []rune
	tmp := 0
	for i := len(a) - 1; i >= 0; i-- {
		if iB > 0 {
			iB--
			if Byte2Number(a[i])+Byte2Number(b[iB])+tmp > 9 {
				r = append(r, rune(Byte2Number(a[i])+Byte2Number(b[iB])-10+tmp))
				tmp = 1

			} else {
				r = append(r, rune(Byte2Number(a[i])+Byte2Number(b[iB])+tmp))
				tmp = 0

			}
		} else {
			if Byte2Number(a[i])+tmp > 9 {
				r = append(r, 0)
				r = append(r, 1)
				tmp = 0

			} else {

				r = append(r, rune(Byte2Number(a[i])+tmp))
				tmp = 0

			}

		}

	}
	if tmp != 0 {
		r = append(r, 1)
	}

	var r2 []rune

	for i := len(r) - 1; i >= 0; i-- {
		r2 = append(r2, r[i]+48)
	}
	s = string(r2)
	return s
}
func Byte2Number(a byte) int {
	for i, c := range base36 {
		if c == rune(a) {
			return i
		}
	}
	return -1
}

func LongFibonacci(n int) string {
	a, b := "0", "1"
	var i int
	for ; i < n; i++ {
		a, b = b, LongAdd(a, b)

	}
	return a
}
func LongAddEx(a, b string, base int) (s string) {
	if len(a) < len(b) {
		a, b = b, a
	}
	iB := len(b)
	var r []rune
	tmp := 0
	for i := len(a) - 1; i >= 0; i-- {
		if iB > 0 {
			iB--
			if Byte2Number(a[i])+Byte2Number(b[iB])+tmp > base-1 {
				r = append(r, rune(Byte2Number(a[i])+Byte2Number(b[iB])-base+tmp))
				tmp = 1

			} else {
				r = append(r, rune(Byte2Number(a[i])+Byte2Number(b[iB])+tmp))
				tmp = 0

			}
		} else {
			if Byte2Number(a[i])+tmp > base-1 {
				r = append(r, 0)
				r = append(r, 1)
				tmp = 0

			} else {

				r = append(r, rune(Byte2Number(a[i])+tmp))
				tmp = 0

			}

		}

	}
	if tmp != 0 {
		r = append(r, 1)
	}

	var r2 []rune

	for i := len(r) - 1; i >= 0; i-- {
		r2 = append(r2, r[i]+48)
	}
	s = string(r2)
	return s
}
func main() {
	// fmt.Println(LongAdd("98765432109876543210987654321098765432109876543210",
	// 	"22222222222222222222222222222222222222222222222222"))
	// fmt.Println(LongFibonacci(500))
	fmt.Println(LongAddEx("55555", "55", 10))
}
