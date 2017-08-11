package ftp

import (
	"ftp"
	"testing"
	"text/scanner"
)

func TestTokenizerScan(t *testing.T) {
	tokenizer := ftp.Tokenizer{}
	tokernizer.Init(`USER cesar \r\n`)
	token := tokenizer.Scan()
	if token != scanner.Ident {
		t.Error("Expecting an identifier")
	}
}
