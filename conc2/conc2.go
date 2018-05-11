package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 100; i++ {
		fmt.Println(<-result)
	}

}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		for {
// 			time.Sleep(time.Millisecond * 500)
// 			c1 <- "Every 500ms"
// 		}
// 	}()

// 	go func() {
// 		for {
// 			time.Sleep(time.Second * 2)
// 			c2 <- "Every 2 seconds"
// 		}
// 	}()
// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }
