package main

import "fmt"

var y = 8

func decrement() int {
	y--
	return y
}

func main() {

	fmt.Println(decrement())
	fmt.Println(decrement())

}
