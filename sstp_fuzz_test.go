// +build gofuzzbeta

package main

import (
	"testing"
)

func FuzzIsSSTP(f *testing.F) {
	f.Add([]byte("x"))
	f.Fuzz(func(t *testing.T, data []byte) {
		IsSSTP(data)
	})
}
