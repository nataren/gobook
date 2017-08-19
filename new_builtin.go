package main

import "fmt"

func main() {
	s := new(struct{})
	a := new([0]int)
	fmt.Printf("s=%p, a=%p\n", s, a)
	fmt.Printf("s=%#v, a=%#v\n", s, a)
}
