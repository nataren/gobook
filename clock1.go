// TCP server that periodically writes the time

package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"strconv"
)

func main() {
	var port = flag.Int("port", 8010, "port where your clock will be listening to")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:" + strconv.Itoa(*port))
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
