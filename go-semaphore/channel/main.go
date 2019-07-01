package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	orders := make([]string, 10)

	orders[0] = "First"
	orders[1] = "Second"
	orders[2] = "Third"
	orders[3] = "Fourth"
	orders[4] = "Fifth"
	orders[5] = "Sixth"
	orders[6] = "Seventh"
	orders[7] = "Eighth"
	orders[8] = "Ninth"
	orders[9] = "Tenth"

	n := runtime.GOMAXPROCS(0)

	fmt.Printf("%v\n", n)
	sem := make(chan string, n)

	for _, order := range orders {
		sem <- order
		go func() {
			print(<-sem)
		}()

	}
	time.Sleep(time.Duration(100) * time.Millisecond)

}

func print(value string) {

	fmt.Printf("%v\n", value)
}
