package functions

import (
	"fmt"
	"math/rand"
	"time"
)

// here this func is returning channel "chan string" with "<-"sign making this returned channel as only receiving type not sending type
func Display(msg string) <-chan string {
	c := make(chan string)
	rand.Seed(time.Now().Unix())
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			t := time.Duration(rand.Intn(1.1e3))
			time.Sleep(t * time.Millisecond)
		}
	}()
	return c
}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
