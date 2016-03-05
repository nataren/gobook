package main

import (
	"os"
	"strings"
	"fmt"
	"log"
	"io"
	"sync"
	"net"
	"bufio"
)

type Clock struct {
	City string
	HostnamePort string
}

func (c Clock) String() string {
	return fmt.Sprintf("%s at %s", c.City, c.HostnamePort)
}

func getClocks(args []string) []Clock {
	clocks := make([]Clock, len(args))
	for i, v := range args {
		values := strings.SplitN(v, "=", 2)
		clocks[i] = Clock {
			City: values[0],
			HostnamePort: values[1],
		}
	}
	return clocks
}

func copy(dst io.Writer, src io.Reader, city string) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			continue
		}
		fmt.Fprintf(dst, fmt.Sprintf("\r%s: %v", city, scanner.Text()))
	}
}

func main() {
	clocks := getClocks(os.Args[1:])
	var wg sync.WaitGroup
	for _, c := range clocks {
		wg.Add(1)
		clock := c
		go func() {
			var conn net.Conn

			// Make sure to clean up after ourselves
			defer func() {
				if conn != nil {
					conn.Close()
				}
				wg.Done()
			}()

			// Connect and echo the put from the connection
			conn, err := net.Dial("tcp", clock.HostnamePort)
			if err != nil {
				log.Printf("error happened while trying to connect to '%s': %v", clock.String(), err)
				return
			}
			copy(os.Stdout, conn, clock.City)
		}()
	}
	wg.Wait()
}
