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
		sum := a[i] + b[i] + car
		str = string(sum%10) + str
		car = sum / 10

	}
	for i := 0; i < len(a); i++ {
		sum := b[i] + car
		str = string(sum % 10)
		car = sum / 10

	}
	return str
}

// прграма печатает смайлики))
func main() {
	fmt.Println(LongAdd("28777779", "20"))
	fmt.Println(Reverse("123"))
}
