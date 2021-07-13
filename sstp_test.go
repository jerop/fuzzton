package main

import (
	"testing"
)

func TestIsSSTP(t *testing.T) {
	var tcs = []struct {
		input string
		want  bool
	}{{
		"seafarer", false,
	}, {
		"ci/cd", false,
	}, {
		"drydock", false,
	}}
	for _, tc := range tcs {
		got := IsSSTP([]byte(tc.input))
		if got != tc.want {
			t.Errorf("got %t want %t", got, tc.want)
		}
	}
}
