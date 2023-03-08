package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in golang - we are going to create a ludo game, dias 1 to 6")

	//rand.Seed(time.Now().UnixNano()) //no need of this, in earlier version we need this also, otherwise the random nos created are repeating, so this is to generate different random nos beased on each nano second,, manual way of setting random value
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough //by using fall through if a case is hit, then the next case will also executed after the current case after that... we can use fallthrough in golang
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spot and roll die again")
		//fallthrough
	default: // default is not mandatory in golang
		fmt.Println("Invalid dice entry")
	}
}
