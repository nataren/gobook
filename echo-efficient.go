// Echo-efficient prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo(parts []string) string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	fmt.Println(echo(os.Args[1:]))
}
