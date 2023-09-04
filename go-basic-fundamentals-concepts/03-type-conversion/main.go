package main

import "fmt"

func main() {
	speed := 100
	force := 2.5
	speed = speed * int(force)
	fmt.Println(speed)

	speed1 := 100
	speed1 = int(float64(speed1) * force)
	fmt.Println(speed1)

	var apple int
	var orange int32
	apple = int(orange)
	fmt.Println(apple)

	num := 65
	alpha := string(num)
	fmt.Println(alpha)

	fmt.Println(string([]byte{104, 105, 65}))

}
