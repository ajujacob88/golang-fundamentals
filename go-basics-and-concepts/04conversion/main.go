package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ref :-  https://youtu.be/3j43y-PFJPI

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate out pizza between 1 and 5")

	readeraju := bufio.NewReader(os.Stdin)

	input, _ := readeraju.ReadString('\n')

	fmt.Println("Thank for rating, ", input)

	//numRating, err := strconv.ParseFloat(input, 64)  ,, because while running, \n also takes and hence conversion not possible, so use the below code,
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //by using the string package we can overcome that..

	if err != nil { //error checking code
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating: ", numRating+1)
	}

}
