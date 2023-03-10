package functions

import (
	"fmt"
	"time"
)

func Booring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}
