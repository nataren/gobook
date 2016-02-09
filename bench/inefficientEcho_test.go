// Echo-inefficient prints its command-line arguments
package gobook

import "testing"

func inefficientEcho(parts []string) string {
	var s, sep string
	for i := 1; i < len(parts); i++ {
		s += sep + parts[i]
		sep = " "
	}
	return s
}

func BenchmarkInefficientEcho(b *testing.B) {
	parts := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	for i := 0; i < b.N; i++ {
		inefficientEcho(parts[:])
	}
}
