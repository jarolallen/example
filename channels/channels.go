package main

import "fmt"

func main() {

	messages := make(chan string) // make a new channel called messages that sends a string

	go func() { messages <- "ping" }() //go function that sends a string into the messages channel

	msg := <-messages //recieve the messages channel into a variable
	fmt.Println(msg)  //print out the variable
}
