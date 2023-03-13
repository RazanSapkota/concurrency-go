package main

import (
	"fmt"
	"time"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	c := functions.Display("john")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Time over")
			return
		}
	}
}
