package main

import "fmt"

type Duration int

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
}
