package main

import "fmt"

func main() {
	fmt.Println("hi")
	var counter int
	counter = 5
	var mask int32
	mask = counter    // type mismatch
}
