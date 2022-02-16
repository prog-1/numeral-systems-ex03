package main

import (
	"fmt"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func LongAdd(a, b string) string {
	if len(a) > len(b) {
		t := a
		a = b
		b = t
	}
	var str string
	var car uint8
	a = Reverse(a)
	b = Reverse(b)
	for i := 0; i < len(a); i++ {
		sum := (a[i] - '0') + (b[i] - '0') + car
		str = string((sum%10 + '0')) + str
		car = sum / 10

	}
	for i := 0; i < len(b)-len(a); i++ {
		sum := (b[i] - '0') + car
		str = string((sum%10)+'0') + str
		car = sum / 10
	}
	if car != 0 {
		str = "1" + str
	}
	//str = Reverse(str)

	return str
}
func main() {
	fmt.Println(LongAdd("55555", "55"))
	fmt.Println(Reverse("123"))
}
