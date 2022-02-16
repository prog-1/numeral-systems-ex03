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
		{"98765432109876543210987654321098765432109876543210", "22222222222222222222222222222222222222222222222222", "120987654332098765433209876543320987654332098765432"},
	} {
		t.Run("", func(t *testing.T) {
			if got := LongAdd(tc.a, tc.b); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
