package main

import "fmt"

type Duration int //here the defined type is Duration and underlying type is int

type gram float64 //here the defined type is gram and underlying type is float64
type ounce float64

func main() {
	var dur Duration = 55040

	fmt.Printf("type of dur is %T \n", dur)

	var dur2 int = 5
	fmt.Printf("typee of dur is %T \n", dur2)

	//result := dur * dur2

	result := dur * Duration(dur2)

	result2 := dur * 5

	result3 := int(dur) * dur2

	fmt.Printf("%v, %v, %v\n", result, result2, result3)

	var g gram = 1000
	var o ounce
	o = ounce(g) * .035274
	fmt.Printf("%g grams is %.2f ounces\n", g, o)
}
