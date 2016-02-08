// Echo prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	//	fmt.Println(os.Args)
	for i, cmd := range os.Args {
		fmt.Printf("%v %s\n", i, cmd)
	}
}
