package main

import "fmt"

func comp(
	f func(x interface{}) interface{},
	g func(y interface{}) interface{},
) func(x interface{}) interface{} {
	return func(x interface{}) interface{} {
		return g(f(x))
	}
}

func main() {
	add5 := comp(
		func(n interface{}) interface{} {
			if nn, ok := n.(int); ok {
				return nn + 2
			}
			return nil
		},
		func(n interface{}) interface{} {
			if nn, ok := n.(int); ok {
				return nn + 3
			}
			return nil
		},
	)
	fmt.Println(add5(1) == 6)
}
