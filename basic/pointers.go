package main

import (
	"flag"
	"fmt"
	"strings"
)

func pointers() {
	println()
	println("******* Pointers *******")
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2, *p is an alias to x
	fmt.Println(x)  // "2"
	var w, y int
	fmt.Println(&w == &w, &w == &y, &w == nil) // "true false false"

	var k = f() // k still in existence after the call of the function
	println(k)

	// each call of f returns a distinct value
	fmt.Println(f() == f()) // "false"

	l := 1
	incr(&l)              // side effect: l is now 2
	fmt.Println(incr(&l)) // "3" (and l is 3)

	flags() // this works if we build the application first. Example of command: bin/basic -s / a b | or bin/basic -help
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}

func flags() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")
	// n and sep are pointers to flag variables, they must be accessed indirectly with *n and *sep
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
