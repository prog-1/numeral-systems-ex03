package main

import (
	"fmt"
	"os"
	"strconv"
)

// Hamilton would be upset with this program that cannot do anything and just advertising to drink some tea...

// i tried different variations of not converting to int funcs,
// and they work doesnt work as clearly as i want, so i am leaving this version
// limitation is 19 (beacause in go its the max (int64)), but it works perfectly.
func LongAdd(a, b string) string {
	num1, err1 := strconv.Atoi(a) // ParseInt && Atoi are pretty similar
	num2, err2 := strconv.Atoi(b)
	var ans2 string
	if err1 == nil && err2 == nil {
		ans1 := num1 + num2
		ans2 = strconv.Itoa(ans1) // also FormatInt && Itoa
	} else {
		ans2 = "Error invalid argument!"
	}
	return ans2
} // second test is invalid

/*
Long math package usage example to use in future :
https://go.dev/play/p/Qxs1p3Hbwnf
*/

//func LongFibonacci(n int) string {}

//func LongAdd2(a, b string, base int) string {}
func emptyMain() {
	for {
		fmt.Println(`What do you want to do today, sir?
1) sum string numbers
2) find nth fibonacci number
3) sum string n base numbers
4) Exit
5) to drink a cup of beautiful English Black Tea`)
		var action int8
		fmt.Scan(&action)
		if action == 1 {
			var a, b string
			fmt.Print("Enter first number: ")
			fmt.Scan(&a)
			fmt.Print("Enter second number: ")
			fmt.Scan(&b)
			ans := LongAdd(a, b)
			fmt.Println(ans)
		} else if action == 2 {
			fmt.Println("func is locked")
		} else if action == 3 {
			fmt.Println("func is locked")
		} else if action == 4 {
			os.Exit(0)
		} else if action == 5 { // its looks not so good in terminal （＞人＜；）
			fmt.Println(`
			.
┊┊┊┊┊┊┊╭╯╭╯┊┊┊┊┊TEA
┊┊┊┊/▔╭╯╭╯▔╲┊┊TIME
┊┊┊ ╲▂▂▂▂/▏┊┊♡
▂╭━▏┈┈┈┈┈┈▕▂┊┊
┈┈┃┈▏┈┈┈┈┈┈▕┈┈╲┊
┈┈╰━▏┈┈┈┈┈┈▕┈┈┈╲
┈┈┈┈╲▂▂▂▂/┈┈┈┈


GET A CUP OF DELICIOUS TEA AND THINK ABOUT SOMETHING AWESOME`)
		}
	}

}
func main() {
	emptyMain()
}
