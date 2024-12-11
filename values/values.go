package main

import "fmt"

func main() {

	// You can "add" strings like in Python
	fmt.Println("go" + "lang")
	// "math" in quotes is a string
	// Without quotes the expression will be evaluated like in Python
	fmt.Println("20+3/5 =", 20+3/5) //floor division
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// true AND false statement is false
	fmt.Println(true && false)
	// true OR false statement is true
	fmt.Println(true || false)
	// NOT true statment is false
	fmt.Println(!true)
}
