package main

import "testing"

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		name      string
		a         string
		b         string
		sumResult string
	}{
		{"simple math", "123", "210", "333"},
		{"Math with memorization ", "55555", "55", "55610"},
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
