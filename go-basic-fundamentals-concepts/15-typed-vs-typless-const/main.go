package main

import "fmt"

func main() {
	const min = 50 //typeless constant

	var num1 int = min     //here min's type automatically became int
	var num2 float64 = min //here min's type automatically became float 64
	var num3 byte = min
	var num4 rune = min
	var num5 uint32 = min

	fmt.Println(num1, num2, num3, num4, num5)

	const max int = 40 //typed constant
	var num6 int = max
	//var num7 float64 = max
	fmt.Println(num6)

	const numb1 = 2
	const pi = 3.14 * numb1

	//default types
	const aju int32 = 25
	jacob := 30 + aju
	fmt.Printf("type of jacob is %T\n", jacob)

	const check = 50
	check2 := 25 + check
	check3 := 25.5 + check
	fmt.Printf("type of check2 is %T\n", check2)
	fmt.Printf("type of check3 is %T\n", check3)
	i := 42
	f := 56.3
	b := true
	s := "hello"
	r := 'A'
	fmt.Printf("type of i is %T\ntype of f is %T\ntype of b is %T\ntype of s is %T\ntype of r is rune or %T\n", i, f, b, s, r)
}
