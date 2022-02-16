package main

import "fmt"

func addZeros(a string, count int) string {
	for i := count; i > 0; i-- {
		a = "0" + a
	}
	return a
}

func LongAdd(a, b string) string {
	var add bool
	if len(a) > len(b) {
		b = addZeros(b, len(a)-len(b))
	} else {
		a = addZeros(a, len(b)-len(a))
	}
	ab := make([]byte, len(a))
	for i := 1; i <= len(a); i++ {
		// 106 or byte(106) is the two digit result of addition, e. g., []byte("4" + "6") = [52 54], 52 + 54 = 106
		if (add && a[len(a)-i]+b[len(b)-i]+byte(1) >= 106) || a[len(a)-i]+b[len(b)-i] >= 106 {
			if add {
				ab = append([]byte{a[len(a)-i] + b[len(b)-i] - byte(58) + byte(1)}, ab...)
			} else {
				ab = append([]byte{a[len(a)-i] + b[len(b)-i] - byte(58)}, ab...)
			}
			if add && len(a)-i == 0 {
				ab = append([]byte("1"), ab...)
			}
			add = true
		} else {
			if add {
				ab = append([]byte{a[len(a)-i] + b[len(b)-i] - byte(48) + byte(1)}, ab...)
				add = false
			} else {
				ab = append([]byte{a[len(a)-i] + b[len(b)-i] - byte(48)}, ab...)
			}
		}
	}
	return string(ab)
}

func main() {
	var a, b string
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	ab := LongAdd(a, b)
	fmt.Println("The sum of numbers: ", ab)
}
