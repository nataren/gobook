package main

import (
	"os"
	"strings"
	"fmt"
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

func (c Clock) String() string {
	return fmt.Sprintf("%s at %s", c.City, c.HostnamePort)
}

type CityTime struct {
	City string
	Time string
}

func (ct CityTime) String() string {
	return fmt.Sprintf("%s: %v", ct.City, ct.Time)
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
			Time: scanner.Text(),
		}
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
				return
			}
			readtime(conn, clock)
		}()
	}
	cityTimes := make([]chan CityTime, len(clocks))
	for i, clock := range clocks {
		cityTimes[i] = clock.Output
	}
	
	// Output to std output
	go func(cityTimes []chan CityTime) {
		wg.Add(1)
		defer wg.Done()
		length := len(cityTimes)
		vals := make([]CityTime, length)
		messages := make([]string, length)
		for {
			for i, citytime := range cityTimes {
				vals[i] = <-citytime
			}
			sort.Sort(ByCity(vals))
			for i, citytime := range vals {
				messages[i] = citytime.String()
			}
			fmt.Printf("\r| %s |", strings.Join(messages, " | "))
		}
	}(cityTimes)
	wg.Wait()
}
