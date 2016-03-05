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
	"sort"
)

type Clock struct {
	City string
	HostnamePort string
	Output chan CityTime
}

type CityTime struct {
	City string
	Time string
}

type ByCity []CityTime

func (bc ByCity) Len() int {
	return len(bc)
}

func (bc ByCity) Less(i, j int) bool {
	return bc[i].City < bc[j].City
}

func (bc ByCity) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}

func (ct CityTime) String() string {
	return fmt.Sprintf("%s: %v", ct.City, ct.Time)
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
			Output: make(chan CityTime),
		}
	}
	return clocks
}

func readtime(src io.Reader, c Clock) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			continue
		}
		c.Output <- CityTime {
			City: c.City,
			Time: fmt.Sprintf("%s: %v", c.City, scanner.Text()),
		}
	}
}

func main() {
	clocks := getClocks(os.Args[1:])
	log.Println("created clocks")
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
			readtime(conn, clock)
		}()
	}
	log.Println("setup time readers")
	cityTimes := make([]chan CityTime, len(clocks))
	for i, c := range clocks {
		cityTimes[i] = c.Output
	}
	
	// Output to std output
	go func(cityCityTimes []chan CityTime) {
		wg.Add(1)
		defer wg.Done()
		length := len(cityCityTimes)
		vals := make([]CityTime, length)
		messages := make([]string, length)
		for {
			for i, citytime := range cityCityTimes {
				vals[i] = <-citytime
			}
			sort.Sort(ByCity(vals))
			for i, citytime := range vals {
				messages[i] = citytime.String()
			}
			fmt.Sprintf("\r%s", strings.Join(messages, "| "))
		}
	}(cityTimes)
	log.Println("setup clockwall output")
	wg.Wait()
}
