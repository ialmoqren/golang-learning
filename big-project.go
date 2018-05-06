package main

import (
	"fmt"
	"time"
)

func main() {

	var first string
	for {
		first = time.Now().Format(time.RFC850)
		fmt.Println(first)
		for {
			if time.Now().Format(time.RFC850) != first {
				break
			}
		}
	}
}
