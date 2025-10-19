package main

import (
	"fmt"
	"math"
)

func root(x float64) float64 {
	return math.Sqrt(x)
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println("The square root of 49 is: ", root(49))
	fmt.Println(swap("Hello", "world!"))
}
