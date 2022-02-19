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
		{"1", "1", "2"},
		{"999", "1001", "2000"},
		{"123", "210", "333"},
		{"55555", "55", "55610"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"120987654332098765433209876543320987654332098765432"},
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
		{10, "55"},
		{500, "139423224561697880139724382870407283950070256587697307264108962948325571622863290691557658876222521294125"},
	} {
		t.Run("", func(t *testing.T) {
			if got := LongFibonacci(tc.n); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
func TestLongAddEx(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		base int
		want string
	}{
		{"", "", 0, ""},
		{"0", "0", 1, "0"},
		{"1", "1", 2, "10"},
		{"999", "1001", 10, "2000"},
		{"AB00", "CD", 16, "55610"},
		{"55555", "55", 10, "55610"},
		{"123", "2", 5, "130"},
	} {
		t.Run("", func(t *testing.T) {
			if got := LongAddEx(tc.a, tc.b, tc.base); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
