//predeclared types

package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		width  uint8 = 255   //range of values that can represenmt uint8 is 0 to 255
		height int16 = 32767 //range of values represent int16 is-32768 to32767,,, lly int32, int64 also having its maximum limit of 32 bit and 64 bit resp.
	)
	fmt.Println(width)
	fmt.Println(width)

	width++ //incrementing  above this value will reset it to the starting value,, similar is the case with float also
	height++

	fmt.Println(width)
	fmt.Println(height)
	fmt.Println("maximum value for int8 type is", math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt)

	var minint int = math.MinInt
	fmt.Println("minimum value for int64 is", minint)
	minint--
	fmt.Println(minint)

}
