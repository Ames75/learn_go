package main

import (
	"fmt"
	"math"
)

func testVar() {
	var name string = "xyz"
	fmt.Println(name)
}

func testMath() {
	fmt.Println(math.Ceil(2.8))
}

func main() {
	// testVar()
	testMath()
	fmt.Printf("hello, world\n")
}
