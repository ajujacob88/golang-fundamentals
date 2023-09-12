package main

import "fmt"

func main() {
	speed := 100
	fmt.Println(speed > 150)

	check := speed == 100

	fmt.Printf("type of check is %T and value is %v\n", check, check)

	var on bool
	fmt.Println(on)
	on = !on
	fmt.Println(on)

	if speed > 90 && on {
		fmt.Println("yes")
	}

	if speed > 90 && !on {
		fmt.Println("no")
	}
}
