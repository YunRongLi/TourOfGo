package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//Basic Variable Types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// byte alias for uint8
// rune alias for int32
//      represents a Unicode code point
// float32 float64
// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var a, b int = 1, 2
	c := 3
	var d, e, f = true, false, "no!"
	fmt.Println(a, b, c, d, e, f)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var g int
	var h float64
	var i bool
	var str string
	fmt.Printf("%v %v %v %q\n", g, h, i, str)

	var x, y int = 3, 4
	f = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	m := 42
	n := 3.142
	o := 0.8670 + 0.5i
	fmt.Printf("m is of type %T\n", m)
	fmt.Printf("n is of type %T\n", n)
	fmt.Printf("q is of type %T\n", o)

	const world = "世界" // constants can't be declared using the := syntax.
	fmt.Println("Hello", world)
	fmt.Printf("%v is of type %T\n", world, world)
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
