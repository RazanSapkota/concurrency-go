package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup

	for i := 0; i < total; i += max {
		limit := max

		if i+max > total {
			limit = total - i
		}

		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("could not dial: %v", err)
				}
				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("could not read from connection: %v", err)
				}
				if string(bs) != "success" {
					log.Fatalf("request error,request: %d", i+j+1)
				}

				fmt.Printf("request %d: success\n", i+j+1)
			}(j)
		}
		wg.Wait()

	}
}
