package main

import (
	"fmt"
	"time"
)

func main() {

	x := 5
	for {
		fmt.Println(x, x, x, time.Now().Format(time.RFC850))
		x += 3
		if x > 25 {
			break
		}
	}
}
