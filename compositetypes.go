package main

import (
	"fmt"
	"regexp"
	"strings"
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

	// regex
	queryArgs := regexp.MustCompile(`\w+=\w*`)
	fmt.Printf("[FindString] %v\n", queryArgs.FindString("https://contoso.com?c=d&x=y"))

	matches := queryArgs.FindAllString("https://contoso.com?a=b&foo=bar", -1)
	if matches == nil {
		fmt.Println("found no matches")
		return
	}
	for _, m := range matches {
		parts := strings.Split(m, "=")
		if len(parts) >= 2 {
			fmt.Printf("[FindAllString] %v->%v\n", parts[0], parts[1])
		}
	}
}
