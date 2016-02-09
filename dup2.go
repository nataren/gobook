package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Info struct {
	Count     int
	Filenames map[string]struct{}
}

func main() {
	counts := make(map[string]Info)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, info := range counts {
		if info.Count > 1 {
			fmt.Printf("%d\t%s\t%s\n", info.Count, line, strings.Join(keys(info.Filenames), " "))
		}
	}
}

func keys(m map[string]struct{}) []string {
	ks := make([]string, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func countLines(f *os.File, counts map[string]Info) {
	filename := f.Name()
	input := bufio.NewScanner(f)
	var e struct{}
	for input.Scan() {
		t := input.Text()
		item, ok := counts[t]
		if ok {
			item.Filenames[filename] = e
			counts[t] = Info{
				Filenames: item.Filenames,
				Count:     item.Count + 1,
			}
		} else {
			counts[t] = Info{
				Filenames: map[string]struct{}{filename: e},
				Count:     1,
			}
		}
	}
}
