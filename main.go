package main

import (
	"fmt"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	c := functions.FanIn(functions.Display("john"), functions.Display("ana"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I am exiting")
}
