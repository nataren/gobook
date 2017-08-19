package main

import (
	"fmt"
	"os"
)

type CustomTypeDefaultValues struct {

	// Nil
	f  func(n int) float64
	ch chan int
	sl []string

	// 0
	n int64

	// Empty
	s string
	m map[complex64]float64
	p *int
}

func main() {
	var ct CustomTypeDefaultValues
	fmt.Printf("%v\n", ct)
	fmt.Printf("len(m)=%d, len(m.sl)=%d\n", len(ct.m), len(ct.sl))

	f, err := os.Open("foo")

	i := 1
	myf := func(n int) {
		fmt.Println("my func: ", n)
	}
	defer myf(i)

	i += 1
	fmt.Println(f, err)
	f, err = os.Create("foo")
	fmt.Println(f, err)

	defer myf(i)
}
