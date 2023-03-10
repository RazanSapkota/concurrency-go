package main

import (
	"fmt"

	"github.com/RazanSapkota/concurrency-go/functions"
)

func main() {
	go functions.Booring("booring")
	fmt.Println("I am listening")
}
