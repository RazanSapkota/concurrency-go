package main

import (
	"fmt"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	c := functions.Booring("booring")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say:%q\n", <-c)
	}
	fmt.Println("I am exiting")
}
