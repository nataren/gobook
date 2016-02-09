// Echo-inefficient prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func echo(parts []string) string {
	var s, sep string
	for i := 0; i < len(parts); i++ {
		s += sep + parts[i]
		sep = " "
	}
	return s
}

func main() {
	fmt.Println(echo(os.Args[1:]))
}
