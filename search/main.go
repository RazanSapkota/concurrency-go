package main

import (
	"fmt"
	"math/rand"
	"time"

	//v3 "github.com/RazanSapkota/concurrency-go/search/v3"
	v4 "github.com/RazanSapkota/concurrency-go/search/v4"
	// v1 "github.com/RazanSapkota/concurrency-go/search/v1"
	// v2 "github.com/RazanSapkota/concurrency-go/search/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := v4.Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
