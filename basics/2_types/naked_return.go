package main

import "fmt"

func split(sum int) (x, y int, x1 string) {
	x = sum * 4 / 9
	y = sum - x
	x1 = "Syed Jafer";
	return
}

func main() {
	fmt.Println(split(17))
}

