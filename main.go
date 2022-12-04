package main

import (
	"fmt"
	"time"
)

func fib(number float64, ch chan float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}
	ch <- x
}

func main() {
	start := time.Now()
	scan := ""
	ch := make(chan float64, 15)
	// quit := make(chan string, 1)
	for i := 1; i < 15; i++ {
		fmt.Scanln(&scan)
		if scan != "quit" {
			go fib(float64(i), ch)
			fmt.Printf("Fib(%v): %v\n", i, <-ch)
		} else {
			panic("quit")
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
