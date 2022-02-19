package main

import (
	"testing"
)

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		base int
		want string
	}{
		{"0", "0", 0, "-1"},
		{"0", "0", -1, "-1"},
		{"55555", "55", 10, "55610"},
		{"123", "2", 5, "130"},
		{"34", "16", 7, "53"},
		{"36", "32", 7, "101"},
		{"36", "63", 7, "132"},
		{"63", "36", 7, "132"},
		{"660", "31", 8, "711"},
		{"31", "660", 8, "711"},
		{"31", "660", 0, "-1"},
		{"31", "660", -1, "-1"},
		{"31", "660", -31, "-1"},
		{"231", "659", 8, "-1"},
		{"1976A3", "3B08", 11, "-1"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := LongAdd(tc.a, tc.b, tc.base); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
