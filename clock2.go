package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port = flag.Int("port", 8000, "the port you want to connect to")
	flag.Parse()
	if port == nil {
		log.Fatal("{port} flag was not specified")
		return
	}
	listener, err := net.Listen("tcp", "localhost"+":"+fmt.Sprintf("%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
