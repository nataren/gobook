// Echo-inefficient prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("it took %E seconds\n", secs)
	fmt.Println(s)
}
