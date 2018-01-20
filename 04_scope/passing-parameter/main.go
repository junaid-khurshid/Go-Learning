package main

import (
	"fmt"

	vis "github.com/juni381/Go-Learning/03_packages/visibility"
)

func main() {

	fmt.Println(vis.Myname)
	fmt.Println(vis.Add(6))
	vis.Print()

}
