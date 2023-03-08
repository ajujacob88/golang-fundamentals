package main

import "fmt"

func main() {
	var p int
	var r float64
	var n float64
	fmt.Println("Enter the Principal Amount")
	fmt.Scan(&p)
	fmt.Println("Enter the rate of Interest")
	fmt.Scan(&r)
	fmt.Println("Enter the no of years")
	fmt.Scan(&n)

	si := (float64(p) * n * r) / 100
	fmt.Printf("The simple interest is %v", si)
}
