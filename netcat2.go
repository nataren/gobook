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
	var port = flag.Int("port", 8010, "port where we want to connect to")
	flag.Parse()
	conn, err := net.Dial("tcp", "localhost:" + strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
