package main

import "fmt"

func main() {
	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = i + 10
	}
	p := &a[3]
	*p = -1
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
