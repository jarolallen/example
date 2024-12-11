package main

import "fmt"

func main() {

	// Decalre a variable and give it a value
	var a = "initial"
	fmt.Println(a)
	// You can decalre multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)
	// Variables can be booleans
	var d = true
	fmt.Println(d)
	// A number declared without a value is considered 0
	var e int
	fmt.Println(e)
	// You can declare a variable with := syntax. Only valid inside functions
	f := "apple"
	fmt.Println(f)
}
