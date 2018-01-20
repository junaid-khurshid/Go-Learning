package main

import "fmt"

func add(x int) int {
	return 42 + x
}

func main() {
	add := add(8)    //variable shadowing
	fmt.Println(add) // add its now result , not function
}
