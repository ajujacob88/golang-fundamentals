package main

import "fmt"

func main() {
	k := 1
	for i := 0; i < 4; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(k, " ")
			k++
		}
		fmt.Println()
	}
}
