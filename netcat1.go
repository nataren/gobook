package main

import (
	"io"
	"log"
	"net"
	"os"
	"flag"
	"strconv"
)

func main() {
	var port = flag.Int("port", 8010, "port where your clock will be listening to")
	flag.Parse()
	conn, err := net.Dial("tcp", "localhost:" + strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
