package main

import "fmt"

func zeros(a string, i int) string {
	for ; i > 0; i-- {
		a = "0" + a
	}
	return a
}

func LongAdd(a, b string) (s string) {
	var transfer byte
	if len(a) > len(b) {
		b = zeros(b, len(a)-len(b))
	} else {
		a = zeros(a, len(b)-len(a))
	}
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + transfer
		transfer = num / 10
		s = string(num%10+'0') + s
	}
	if transfer > 0 {
		s = "1" + s
	}
	return s
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
1) Sum of two numbers
2) Fibonacci
3) Quit`)
	fmt.Scanln(&choice)
	return choice
}

func main() {
	var a, b string
	var n int
	for {
		choice := mainMenu()
		if choice == 1 {
			fmt.Print("Enter first number:")
			fmt.Scanln(&a)
			fmt.Println("Enter second number:")
			fmt.Scanln(&b)
			fmt.Printf("%v + %v is %v", a, b, LongAdd(a, b))
		} else if choice == 2 {
			fmt.Print("Enter the  N-th Fibonacci sequence number:")
			fmt.Scanln(&n)
			fmt.Printf("%vth Fibonacci number is %v", n, LongFibonacci(n))
		} else if choice == 3 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
