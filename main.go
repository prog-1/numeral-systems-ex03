package main

import "fmt"

func LongAdd(a, b string) (s string) {
	if len(a) < len(b) {
		a, b = b, a
	}
	iB := len(b)
	var r []rune
	tmp := 0
	for i := len(a) - 1; i >= 0; i-- {
		if iB > 0 {
			fmt.Println(iB)
			iB--
			fmt.Println(a[i], "  ", b[iB])
			if a[i]+b[iB] > 105 {
				r = append(r, rune(a[i]+b[iB]-58+byte(tmp)))
				tmp = 1
			} else {
				r = append(r, rune(a[i]+b[iB]-48+byte(tmp)))
				tmp = 0
			}
		} else {
			r = append(r, rune(a[i]+byte(tmp)))
			tmp = 0
		}
		fmt.Println(r)
		//I'm have no idea how to use runes instead of ints here...
	}
	if tmp != 0 {
		r = append(r, '1')
	}
	var r2 []rune
	for i := len(r) - 1; i >= 0; i-- {
		r2 = append(r2, r[i])
	}
	s = string(r2)
	return s
}
func main() {
	fmt.Println(LongAdd("", ""))
}
