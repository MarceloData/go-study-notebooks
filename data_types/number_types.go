package main

import (
	"fmt"
	"math"
)

func main() {
	// surface()
	// mandelbrot_main()
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges -> different types
	var compote = int(apples) + int(oranges)
	fmt.Print(compote)

	f := 3.141
	i1 := int(f)
	fmt.Println(f, i1)
	f = 1.99
	fmt.Println(int(f))

	f1 := 1e100
	i2 := int(f1)
	fmt.Println(i2)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x1 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x1)

	ascii := 'a'
	unicode := 'D'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	var f2 float32 = 16777216
	fmt.Println(f2 == f2+1)
	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	fmt.Println(e, Avogadro, Planck)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

	var x3 complex128 = complex(1, 2)
	var y3 complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x3 * y3))
	fmt.Println(imag(x3 * y3))

	fmt.Println(1i * 1i)

	x4 := 1 + 2i
	y4 := 3 + 4i
	fmt.Println(x4, y4)

	var h string

	fmt.Println(h != "" && h[0] == 'x')

	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	c := s[len(s)]
	fmt.Println(c)
}
