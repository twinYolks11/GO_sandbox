package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func root(x float64) float64 {
	return math.Sqrt(x)
}

func swap(x, y string) (string, string) {
	return y, x
}

func looping() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	for sum < 1000 {
		sum += sum
	}
	sum = 0
	for sum < 1000 {
		sum += sum
	}

	for {
		// infinite loop
		break
	}

	fmt.Println(sum)
}

func conditionals() (number float64) {
	x, y := 24.00, 132.00
	if x < 0 {
		return math.Sqrt(-x) + y
	}

	if v := math.Pow(y, x); v < 10000 {
		return v
	} else {
		return math.Sqrt(x) - y
	}
}

func definiteSqrt(x float64) (theRoot float64) {
	theRoot = 1.0

	for i := 0; i < 10; i++ {
		theRoot -= (theRoot*theRoot - x) / (2 * theRoot)
		fmt.Println(theRoot)
	}
	return theRoot
}

func indefiniteSqrt(x float64) (theRoot float64) {
	theRoot = 1.0
	iterations := 0
	for {
		original := theRoot
		theRoot -= (theRoot*theRoot - x) / (2 * theRoot)
		//fmt.Println(theRoot)
		iterations++
		if math.Abs(original-theRoot) < 0.00005 {
			print("number of iterations: ", iterations)
			return theRoot
		}
	}
}

func switches() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	t := time.Now()
	// evaluates to swtich true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func magic() {
	// deferred function calls are pushed onto a stack and executed in LIFO order
	defer fmt.Println("world")
	fmt.Println("hello ")
}

// Basic types are bool, string, int, int8..., uint, uint8..., uintptr, byte (alias for uint8)
// rune (alias for int32) float32, float64, complex64, complex128

// variables that are uinintialized are given their 'zero' value: 0, false, ""

func main() {
	i, j, what, is, this := 49, 89, true, "this is pretty", "cool"
	fmt.Println("Hello world!")
	fmt.Println("The square root of 49 is: ", root(float64(i)))
	fmt.Println(swap("Hello", "world!"))
	fmt.Println(i, j, what, is, this)
	definiteSqrt(float64(j))
	indefiniteSqrt(float64(j))
	switches()
	magic()
}
