package main

import "testing"

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"empty", "", "", ""},
		{"test with 0", "0", "0", "0"},
		{"simple", "1", "2", "3"},
		{"same even numbers", "2", "2", "4"},
		{"same odd numbers", "7", "7", "14"},
		{"from README", "123", "210", "333"},
		{"from README", "55555", "55", "55610"},
		{"from README", "98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"120987654332098765433209876543320987654332098765432"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongAdd(tc.a, tc.b); got != tc.want {
				t.Errorf("LongAdd(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
