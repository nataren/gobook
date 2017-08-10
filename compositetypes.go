package main

import (
	"fmt"
	"regexp"
)

type Foo struct {
	field1 string
}

func (foo *Foo) Amethod(arg1, arg2 int) {
	fmt.Println("Invoking Amethod with args:", arg1, arg2)
	fmt.Printf("And `field1`=%v\n", foo.field1)
	foo.field1 = "set by Amethod"
}

func (foo Foo) Bmethod(arg1, arg2 float32) {
	fmt.Println("Invoking Bmethod with args:", arg1, arg2)
	fmt.Printf("And `field1`=%v\n", foo.field1)
	foo.field1 = "set by Bmethod"
}

func main() {
	foo := Foo{field1: "my field"}
	fmt.Println(foo)

	// regex
	validID := regexp.MustCompile(`foo[0-9]+`)
	fmt.Println(validID.MatchString("foo1"))
}
