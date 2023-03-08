package main

import "fmt"

func main() {
	//ref: https://youtu.be/JoUSa8jtadE

	fmt.Println("Welcome to Array in Go")

	var fruitList [7]string //the no of items is must, here 7 items space is allloted, if we not enter the element in a space, that location will be treated as a space,

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[6] = "Peach" //here we omited positions 2 to 5, so in output that locations is automaticaaly given a space. so care should be taken while telling the required no of items.

	fmt.Println("Fruit List is: ", fruitList)

	fmt.Println("Length of Fruit List is: ", len(fruitList))

	var vegList = [6]string{"potato", "beans", "mushroom"} //string or int or float, anything is possible.
	fmt.Println("veg list is: ", vegList)
	fmt.Println("length of veg list is: ", len(vegList))

}
