package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine") //functions with go can be run concurrently

	go func(msg string) { //functions with go can be run concurrently
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) //waits for a second
	fmt.Println("done")
}
