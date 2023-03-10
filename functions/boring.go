package functions

import (
	"fmt"
	"math/rand"
	"time"
)

// here this func is returning channel "chan string" with "<-"sign making this returned channel as only receiving type not sending type
func Booring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
