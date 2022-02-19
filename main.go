package main

import "fmt"

func addZeros(a string, i int) string {
	for ; i > 0; i-- {
		a = "0" + a
	}
	return a
}

func LongAdd(a, b string) (res string) {
	if len(a) > len(b) {
		a, b = b, a
	}
	var carry byte
	a = addZeros(a, len(b)-len(a))
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + carry
		carry = num / 10
		res = string(num%10+'0') + res
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

func LongFibonacci(n int) (s string) {
	a, b := "0", "1"
	for i := 0; i < n; i++ {
		a, b = b, LongAdd(a, b)
	}
	return a
}

func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) Function should
return the sum of a and b.
2) Program that outputs N-th Fibonacci sequence number.
3) Quit.`)
	fmt.Scanln(&choice)
	return choice
}

func main() {
	choice := mainMenu()
	if choice == 1 {
		fmt.Println("The program accepts two decimal numbers encoded as strings and return sum of a and b")
		var a, b string
		fmt.Println("Enter a:")
		fmt.Scan(&a)
		fmt.Println("Enter b:")
		fmt.Scan(&b)
		fmt.Println("The sum of a and b is", LongAdd(a, b))
	}
	if choice == 2 {
		var nth int
		fmt.Println("Enter n-th term")
		fmt.Scan(&nth)
		fmt.Println("The n=th term of Fibonacci is:", LongFibonacci(nth))
	}
	if choice == 3 {
		return
	}
}
