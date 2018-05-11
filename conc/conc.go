package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}
	wg.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

	// go foo(fooVal, 5)
	// go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal
	// fmt.Println(v1, v2)
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func cleanup() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered in cleanup:", r)
// 	}
// 	wg.Done()
// }

// func say(s string) {
// 	defer cleanup()
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 		if i == 2 {
// 			panic("Oh, i is equal to 2")
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go say("Hey")
// 	wg.Add(1)
// 	go say("there")
// 	wg.Wait()
// }
