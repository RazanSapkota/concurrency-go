package main

import (
	"fmt"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	c := make(chan string)
	go functions.Booring("booring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say:%q\n", <-c)
	}
	close(c)
	fmt.Println("I am exiting")
}
