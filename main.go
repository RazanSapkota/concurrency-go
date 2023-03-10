package main

import (
	"fmt"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	john := functions.Booring("john")
	ana := functions.Booring("ana")

	for i := 0; i < 5; i++ {
		fmt.Println(<-john)
		fmt.Println(<-ana)
	}
	fmt.Println("I am exiting")
}
