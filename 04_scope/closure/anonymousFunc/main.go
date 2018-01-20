package main

import "fmt"

func main() {

	var x = 4

	increment := func() int { // naming inside the func for func can be done like this
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

//anonymous function defined for variable assigning
