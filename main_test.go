package main

import "testing"

func TestLongAdd(t *testing.T) {
	for _, tc := range []struct {
		testn int
		a     string
		b     string
		want  string
	}{
		{1, "", "", ""},
		{2, "0", "0", "0"},
		{3, "1", "", "1"},
		{4, "", "1", "1"},
		{5, "1", "1", "2"},
		{6, "123", "321", "444"},
		{7, "123", "210", "333"},
		{8, "55555", "55", "55610"},
		{9, "98765432109876543210987654321098765432109876543210",
			"22222222222222222222222222222222222222222222222222",
			"120987654332098765433209876543320987654332098765432"},
	} {
		if got := LongAdd(tc.a, tc.b); got != tc.want {
			t.Errorf("Test number %v - error: got = %v, want = %v", tc.testn, got, tc.want)
		}

	}
}
