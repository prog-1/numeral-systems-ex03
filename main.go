package main

// I know there is a better way of doing this

func addZeros(a string, time int) string {
	for i := time; i > 0; i-- {
		a = "0" + a
	}
	return a
}

func Normalize(a, b string) (string, string) {
	if len(a) > len(b) {
		b = addZeros(b, len(a)-len(b))
	} else {
		a = addZeros(a, len(b)-len(a))
	}
	return a, b
}

func LongAdd(a, b string) string {
	result := ""
	var add uint8
	a, b = Normalize(a, b)
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + add
		if num > 9 {
			add = 1
		} else {
			add = 0
		}
		result = string((num%10)+'0') + result
	}
	if add != 0 {
		result = "1" + result
	}
	return result
}
func LongFibonacci(n int) string {
	prev, now := "0", "1"
	for i := 1; i < n; i++ {
		prev, now = now, LongAdd(prev, now)
	}
	return now
}

func main() {

}
