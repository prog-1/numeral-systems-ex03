package main

import "fmt"

func LongAdd(a, b string) string {
	...
}

func main() {
	fmt.Println("The program accepts two decimal numbers encoded as strings and return sum of a and b")
	var a, b string
	fmt.Println("Enter a:")
	fmt.Scan(&a)
	fmt.Println("Enter b:")
	fmt.Scan(&b)
	fmt.Println("The sum of a and b is", LongAdd(a, b))
}
