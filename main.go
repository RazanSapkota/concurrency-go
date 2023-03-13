package main

import (
	"fmt"
	"time"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	c := functions.Display("john")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Time over")
			return
		}
	}
}
