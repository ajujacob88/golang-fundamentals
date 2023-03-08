package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the input char:")
	var character rune
	_, err := fmt.Scanf("%c", &character)
	if err != nil {
		fmt.Println("Error check")
	}
	fmt.Printf("The character entered is %c", character)

}
