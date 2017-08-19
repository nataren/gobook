package main

import "fmt"

type CustomTypeDefaultValues struct {

	// Nil
	f  func(n int) float64
	ch chan int
	// TODO: slice

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
	fmt.Printf("len(m)=%d\n", len(ct.m))
}
