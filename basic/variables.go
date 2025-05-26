package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

var global *int

// x scape from f1, so it will be allocated in the heap memory,
// (this is a point of attention to avoid the consuption of resources unnecessarily)
func fglobal() {
	var x int
	x = 1
	global = &x
}

// y do not scape from g, so it will be placed in the stack memory
func g() {
	y := new(int)
	*y = 1
	fmt.Println(*y)
}

func variables() {

	// Basic variables declaration
	var s string // basic var declaration: default value of string is ""
	fmt.Println(s)
	fmt.Println("****** Basic var declarations *******")
	var i, j, k int                 // int,int,int
	var b, f, d = true, 2.3, "four" // bool, float64, string
	fmt.Println(i, j, k)
	fmt.Println(b, f, d)
	var name string
	var file, err = os.Open(name) // os.Open returns a file and a error
	fmt.Println(err)
	fmt.Println(file)
	file.Close()

	// Short variable declarations
	// short declarations must set at least one new variable, or will return a compile error
	fmt.Println()
	fmt.Println("****** Short var declarations *******")
	t := 0.0
	freq := rand.Float64() * 3.0
	i1, j1 := 0, 1
	f1, err1 := os.Open(name)
	println(t, freq, i1, j1, f1)
	println(err1)
	f1.Close()

	println()
	println("******* The 'new' function *******")
	p1 := new(int)   // p1, of type *int, points to an unnamed int variable
	fmt.Println(*p1) // "0"
	*p1 = 2          // sets the unnamed int to 2
	fmt.Println(*p1) // "2"
	//each call to new returns a distinct variable with a unique address
	p2 := new(int)
	q2 := new(int)
	fmt.Println(p2 == q2) // "false"
	fmt.Println(newInt())
	fmt.Println(newInt1())

	// Lifetime of variables
	println()
	println("******* Lifetime of variables *******")
	fglobal()
	fmt.Println(*global)
	g()
	// fmt.Println(y), y does not exist anymore

	println()
	println("******* Assignments *******")
	x := 1 // named variable
	var y bool
	p := &y
	*p = true // indirect variable
	type Person struct {
		name string
	}
	var person Person
	person.name = "bob" // struct field
	var count [2]int
	scale := 2
	count[x] = count[x] * scale // array or slice or map element
	fmt.Println(x, *p, person, count)

	println()
	println("******* Assignment operators *******")
	count[x] *= scale
	v := 1
	fmt.Println(v)
	v++ // same as v = v + 1, v becomes 2
	fmt.Println(v)
	v-- // same as v = v - 1; v becomes 1 again
	fmt.Println(v)

	println()
	println("******* Tuple Assignment *******")
	x1 := 1
	y1 := 2
	fmt.Println(x1, y1)
	x1, y1 = y1, x1
	fmt.Println(x1, y1)
	var a = [2]int{10, 20}
	fmt.Println(a[0], a[1])
	a[0], a[1] = a[1], a[0]
	fmt.Println(a[0], a[1])
	i, j, k = 2, 3, 5
	fmt.Println(i, j, k)
	f2, err2 := os.Open("foo.txt")
	fmt.Println(f2)
	fmt.Println(err2)
	_, err3 := os.Open("foo.txt") // blank identifier, if we do not want to use one of the variables returned by the function
	fmt.Println(err3)
}

// Identical behavior of new
func newInt() *int {
	return new(int)
}

func newInt1() *int {
	var dummy int
	return &dummy
}
