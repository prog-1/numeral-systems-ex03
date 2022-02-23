package main

import "testing"

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		{"0", "0", "0"},
		{"135", "0", "135"},
		{"123", "210", "333"},
		{"999", "1001", "2000"},
		{"55555", "55", "55610"},
		{"98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"120987654332098765433209876543320987654332098765432"},
	} {
		got := LongAdd(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("ERR: LongAdd(%s, %s): got: %s, want: %s", tc.a, tc.b, got, tc.want)
		}
	}
}
