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

func memoize(f func(x interface{}) interface{}) func(x interface{}) interface{} {
	m := make(map[interface{}]interface{})
	return func(x interface{}) interface{} {
		if k, ok := x.(int); ok {
			if v, ok := m[k]; ok {
				fmt.Println("retrieving for cache: ", k)
				return v
			} else {
				fmt.Println("computing for: ", k)
				v = f(k)
				m[k] = v
				return v
			}
		}
		return nil
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
	mf := memoize(func(i interface{}) interface{} {
		if n, ok := i.(int); ok {
			return n + 5
		}
		return nil
	})
	fmt.Println(mf(1))
	fmt.Println(mf(1))
	fmt.Println(mf(1))
	fmt.Println(mf(1))
}
