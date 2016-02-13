package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const PREFIX = "http://"

	for _, tentativeURL := range os.Args[1:] {
		var url string
		if strings.HasPrefix(tentativeURL, PREFIX) {
			url = tentativeURL
		} else {
			url = PREFIX + tentativeURL
		}
		resp, err := http.Get(url)
		showStatus(resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func showStatus(resp *http.Response) {
	if resp != nil {
		fmt.Printf("\nstatus code: %v\n", resp.Status)
	}
}
