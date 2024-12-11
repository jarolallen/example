package main

func example() { //should have a return type set to float32 / float64
	var str string = "" //strings don't have a boolean true / false value
	if str {
		// this will not run since str is not a boolean
		x := 5 / 2 // x is not used for anything
	}
	return x // x does not exist outside of the function it was declared in
}

func main() {

}
