package main

import "fmt"

func main() {
	var x, y int
	var p *int
	fmt.Println(&x == &y, &x == &x, &x == nil, p == nil)
}
