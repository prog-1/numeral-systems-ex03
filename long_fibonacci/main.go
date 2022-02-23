package main

import "fmt"

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

func LongAdd(a, b string) string {
	var s string
	var add uint8 // uint8 = byte
	a, b = MakeNormal(a, b)

	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + add // '0' = 48
		add = num / 10
		s = string((num%10)+'0') + s // '0' = 48
	}
	if add != 0 {
		s = "1" + s
	}
	return s
}

func LongFibonacci(n int) string {
	var a, b string
	if n >= 0 {
		a, b = "0", "1"
		for i := 0; i < n; i++ {
			a, b = b, LongAdd(a, b)
		}
	} else {
		a = "-1"
	}
	return a
}

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	if LongFibonacci(n) != "-1" {
		fmt.Printf("The %d-th Fibonacci sequence number: %v\n", n, LongFibonacci(n))
	} else {
		fmt.Println("ERROR: Incorrect n")
	}
}
