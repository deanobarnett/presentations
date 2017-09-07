package main

import "testing"

func Reverse(s string) {
	// Reverse me now
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Reverse(10)
	}
}
