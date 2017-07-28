package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type Clock struct {
	City         string
	HostnamePort string
	Output       chan CityTime
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
		clocks[i] = Clock{
			City:         values[0],
			HostnamePort: values[1],
			Output:       make(chan CityTime),
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
		c.Output <- CityTime{
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
		go func(city string) {
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
		}(c.City)
	}
	cityTimes := make([]chan CityTime, len(clocks))
	for i, clock := range clocks {
		cityTimes[i] = clock.Output
	}

	// Output to std output
	go func(cityTimes []chan CityTime) {
		wg.Add(1)
		defer wg.Done()
		for {
			var citiesAndTimes []CityTime
			var messages []string
			for _, citytime := range cityTimes {
				select {
				case v := <-citytime:
					citiesAndTimes = append(citiesAndTimes, v)
				default:
				}
			}
			sort.Sort(ByCity(citiesAndTimes))
			for _, citytime := range citiesAndTimes {
				messages = append(messages, citytime.String())
			}
			fmt.Printf("\r| %s |", strings.Join(messages, " | "))
			time.Sleep(999 * time.Millisecond)
		}
	}(cityTimes)
	wg.Wait()
}
