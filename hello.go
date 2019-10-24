package main

import (
	"fmt"
	"github.com/ilyhacker/hello/computer"
)

func main() {
	fmt.Printf("Hello\n")
	fmt.Printf("%d\n", computer.Add(1, 2))
	fmt.Printf("%d\n", computer.Sub(1, 2))
}
