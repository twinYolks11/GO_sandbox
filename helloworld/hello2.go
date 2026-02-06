package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func pointers() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func structs() {
	v := Vertex{1, 2}
	fmt.Print(v.X)
	p := &v
	p.X = 1e9
	fmt.Println("pointer struct: ", v)
	// v1 := Vertex{1, 2}  // type Vertex
	// v2 := Vertex{X: 1}  // Y:0 is implicit
	// v3 := Vertext{}     // X:0 and Y:0
	// p1 := &Vertex{1, 2} // has type *Vertex
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slices do not store any data, just references the original array
// changing an element in a slice changes the element in its underlying array
// slices have a length and a capacity:
// length = number of elments is contains
// capacity = number of elements in underlying array
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // half open, includes elements 1,2, and 3
	fmt.Println(s)

	// Slice literal
	// creates the array and the slice of the array
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	cool := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
		{7, true},
	}
	fmt.Println(cool)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func makingSlices() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

// if the underlying array is not large enough, a new one will be created
func appending() {
	var s []int
	printSlice("s", s)

	s = append(s, 0)
	printSlice("s", s)

	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2, 3, 4)
	printSlice("s", s)
}

func ranges() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

type Id struct {
	name string
	id   int
}

var m map[string]Id

func maps() {
	m = make(map[string]Id)
	m["Original"] = Id{"Peter Parker", 1}
	fmt.Println(m["Original"])

	// Map literal
	// var s = map[string]Id{
	// 	"Hidden": Id{"Spiderman", 2},
	// }
	// or
	// var y = map[string]Id{
	// 	"wow": {"Owen Wilson", 5},
	// }

	var k = make(map[string]int)

	//Insert or update
	k["key"] = 42

	//retrieve
	elem := k["key"]

	//delete
	delete(k, "key")

	//Check if a key is present
	elem, ok := k["key"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("Key not found")
	}
}

func wordCount(s string) map[string]int {
	var counts = make(map[string]int)
	var fields = strings.Fields(s)

	for word := range fields {
		counts[fields[word]] += 1
	}
	return counts
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functions() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func closures() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pointers()
	// structs()
	// arrays()
	// slices()
	// makingSlices()
	// appending()
	// ranges()
	maps()
	fmt.Println(wordCount("a skunk sat on a stump. the skunk said the stump stunk. the stump said the skunk stunk. therefore they both stunk."))
	functions()
	pos, neg := closures(), closures()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
