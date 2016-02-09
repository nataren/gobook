// Echo-efficient prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	result := strings.Join(os.Args[1:], " ")
	secs := time.Since(start).Seconds()
	fmt.Printf("it took %E seconds\n", secs)

	fmt.Println(result)
}
