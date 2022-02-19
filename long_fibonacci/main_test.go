package main

import (
	"testing"
)

func TestLongFibonacci(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{-1, "-1"},
		{2, "1"},
		{3, "2"},
		{4, "3"},
		{5, "5"},
		{10, "55"},
		{11, "89"},
		{12, "144"},
		{20, "6765"},
		{21, "10946"},
		{26, "121393"},
		{28, "317811"},
		{-28, "-1"},
		{500, "139423224561697880139724382870407283950070256587697307264108962948325571622863290691557658876222521294125"},
		{-500, "-1"},
		{-190214513459040729, "-1"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := LongFibonacci(tc.n); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
