package main

import "fmt"

func LongAdd(a, b string) (s string) {
	return s
}
func LongFibonacci(n int) (s string) {
	return s
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
			fmt.Printf("%vth Fibonacci number is %v", a, LongFibonacci(n))
		} else if choice == 3 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
