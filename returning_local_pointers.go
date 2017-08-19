package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	var p1 = f()
	var p2 = f()
	fmt.Printf("p1==p2=%t\n", p1 == p2)
	fmt.Println(p1, p2, *p1, *p2)

	var p3 = f()
	var p4 = f()
	*p3 = 2
	fmt.Println(p3, p4, *p3, *p4)
	fmt.Println(*p3 + *p4)
}
