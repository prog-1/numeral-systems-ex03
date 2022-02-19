package main

import "testing"

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		{"", "", ""},
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"0", "12", "12"},
		{"123", "210", "333"},
		{"999999", "11", "1000010"},
		{"55555", "55", "55610"},
		{"55", "55555", "55610"},
		{"98765432109876543210987654321098765432109876543210", "22222222222222222222222222222222222222222222222222", "120987654332098765433209876543320987654332098765432"},
	} {
		t.Run("", func(t *testing.T) {
			if got := LongAdd(tc.a, tc.b); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestLongFibonacci(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "1"},
		{5, "5"},
		{10, "55"},
		{13, "233"},
		{26, "121393"},
		{50, "12586269025"},
		{500, "139423224561697880139724382870407283950070256587697307264108962948325571622863290691557658876222521294125"},
	} {
		t.Run("", func(t *testing.T) {
			if got := LongFibonacci(tc.n); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
