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

}
