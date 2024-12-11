package main //Every program in go is part of a package. The program belongs to the main package

import "fmt" //Lets us import from the format package that inclues a print output function

//There's a main fuction similar to C. Everything within the main function is ran when
func main() {
	//Indenation is only for readability

	fmt.Println("hello world") //Hitting Enter implicitly adds a ; to end a line. Semicolon is required in C but not Python.
}
