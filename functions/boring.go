package functions

import (
	"fmt"
	"time"
)

func Booring(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
