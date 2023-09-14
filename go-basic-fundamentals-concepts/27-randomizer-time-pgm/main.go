package main

import (
	"fmt"
	"math/rand"
	"time"
)

//seed the randomizer with time

func main() {
	t := time.Now().UnixNano()

	fmt.Println(t)

	rand.Seed(t)

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	max := 25
	//use the below instead of depreciated rand.seed
	// Seed your custom random generator with a specific seed value.
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Use your custom random generator to generate random numbers.
	randomNumber := myRand.Intn(max) + 1
	fmt.Println("Random Number:", randomNumber)

}
