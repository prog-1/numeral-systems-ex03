package main

import "testing"

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		name      string
		a         string
		b         string
		sumResult string
	}{

		{"zero", "0", "0", "0"},
		{"simple math", "123", "210", "333"},
		{"Math with memorization ", "55555", "55", "55610"},
		{"Nines", "99999", "1", "100000"},
		{"Long number", "98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"120987654332098765433209876543320987654332098765432"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongAdd(tc.a, tc.b); got != tc.sumResult {
				t.Errorf("LongAdd(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.sumResult)
			}
		})
	}
}

func TestLongAddWithBase(t *testing.T) {
	for _, tc := range []struct {
		name      string
		a         string
		b         string
		base      int
		sumResult string
	}{
		{"base 10 sum", "55555", "55", 10, "55610"},
		{"base 5 sum", "123", "2", 5, "130"},
		{"base 2 sum", "1", "1", 2, "10"},
		{"base 7 sum", "66666", "1", 7, "100000"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongAddWithBase(tc.a, tc.b, tc.base); got != tc.sumResult {
				t.Errorf("LongAdd(%v, %v, %v) = %v, want %v", tc.a, tc.b, tc.base, got, tc.sumResult)
			}
		})
	}
}

func TestLongFibonacci(t *testing.T) {
	for _, tc := range []struct {
		n      int
		Result string
	}{
		{0, "0"},
		{1, "1"},
		{10, "55"},
		{500, "139423224561697880139724382870407283950070256587697307264108962948325571622863290691557658876222521294125"},
	} {
		t.Run("Test", func(t *testing.T) {
			if got := LongFibonacci(tc.n); got != tc.Result {
				t.Errorf("LongFibonacci(%v) = %v, want %v", tc.n, got, tc.Result)
			}
		})
	}
}
