package main

import "fmt"

func main() {
	//v := 42 // change me!
	v := 'a'
	fmt.Printf("v is of type %T\n", v)
}

//But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

//i := 42           // int
//f := 3.142        // float64
//g := 0.867 + 0.5i // complex128

