package ftp

import (
	"fmt"
	"strings"
	"text/scanner"
)

type Tokenizer struct {
	s scanner.Scanner
}

func (tokenizer *Tokenizer) Init(s string) {
	s.Init(strings.NewReader(s))
}

func (tokenizer *Tokenizer) Scan() rune {
	return s.Scan()
}

func main() {
	var tokenizer Tokenizer
	var token rune

	tokenizer.Init(`USER cesar\r\n`)
	for token != scanner.EOF {
		token = tokernizer.Scan()
		fmt.Println(token)
	}
}
