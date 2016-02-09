// Echo-efficient prints its command-line arguments

package gobook

import (
	"strings"
	"testing"
)

func efficientEcho(parts []string) string {
	return strings.Join(parts, " ")
}

func BenchmarkEfficientEcho(b *testing.B) {
	parts := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	for i := 0; i < b.N; i++ {
		efficientEcho(parts[:])
	}
}
