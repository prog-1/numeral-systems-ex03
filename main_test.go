package main

import (
	"reflect"
	"testing"
)

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		//{"98765432109876543210987654321098765432109876543210",
		//"22222222222222222222222222222222222222222222222222",
		//"120987654332098765433209876543320987654332098765432"},
		{"123", "210", "333"},
		{"55555", "55", "55610"},
		{"1111111111111111111", "1111111111111111111", "2222222222222222222"},
		{"0", "0", "0"},
		{"", "", "Error invalid argument!"},
		{"absab", "3242", "Error invalid argument!"},
		{"1", "1", "2"},
		{"hjedk", "jfrbga123", "Error invalid argument!"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := LongAdd(tc.a, tc.b); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
