package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getURL(tentativeURL string) string {
	const PREFIX = "http://"
	if strings.HasPrefix(tentativeURL, PREFIX) {
		return tentativeURL
	}
	return PREFIX + tentativeURL
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		fetchall(i, dir)
	}
}

func fetchall(i int, base string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(getURL(url), ch, i, base)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, i int, base string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	writeErr := writeFile(base, url, i, b)
	if writeErr != nil {
		fmt.Printf("error writing file: %v\n", err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, len(b), url)
}

func writeFile(basedir string, uri string, i int, bytes []byte) error {
	u, err := url.Parse(uri)
	f, err := os.Create(filepath.Clean(filepath.Join(basedir, fmt.Sprintf("%v-%s", i, u.Host))))
	check(err)
	defer f.Close()
	_, writeErr := f.Write(bytes)
	if writeErr != nil {
		f.Sync()
	}
	return writeErr
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
