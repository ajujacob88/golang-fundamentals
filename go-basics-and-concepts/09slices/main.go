package main

import (
	"fmt"
	"sort"
)

func main() {
	//ref: https://youtu.be/k7hVj8QL9Co

	fmt.Println("Welcome to slices in Go")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitless is %T \n", fruitList)

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	//below thing is very important in dealing with databases.
	//the column syntax is used to slice up your slice
	fruitList = append(fruitList[2:]) //this will remove the 0th and 1st position
	fmt.Println(fruitList)

	var vegList = []string{"Carrot", "Beans", "Beetroot", "Tomato", "cauliflower", "pappaya", "ladiesfinger"} //this is a slice[] holding values of string

	fmt.Println(vegList)

	vegList = append(vegList[:2]) // this will remove elements from position 2 onwards
	fmt.Println(vegList)
	fmt.Println("Now 1:5")
	vegList = append(vegList[1:5]) //the last range(5th position) is always no inclusive and the first range(1st position) is inclusive
	fmt.Println(vegList)

	//for memory allocation 2 keywords are there, new keyword and make keyword
	highScores := make([]int, 4) //this is slice
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 //this will gives erro, because 4 values are allowed(we already set input as 4)

	highScores = append(highScores, 555, 666, 321) //as soon as we are using the method append, go is reallocating the memory and all of the values are accomodated.
	fmt.Println(highScores)

	sort.Ints(highScores) //this sorts in ascending order  // these all are availablle in sliced, not in the array
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) // this is a boolean values, asking questions like whether it is sorted or not ?

	// REMOVING VALUES FROM SLICE BASED ON INDEX
	// THIS IS VERY IMP TOPIC 7 USEFUL IN DATABASE AND OTHER OPERATIONS - REF: https://youtu.be/931nR5TGCAk

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby", "Dart", "Flutter"}
	fmt.Println(courses)
	var index int = 2 // aim is to remove the second element which is swift
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}

/*
The Go programming language specification does not guarantee any specific order of iteration when ranging over a map.
Therefore, the order of the printed output may vary between different Go compilers or even between different runs of the same code.

However, in some Go compilers, including the official Go compiler and some third-party compilers, the implementation of the map type iterates over the keys in a sorted order by default.
This means that when you print the map, the keys will appear in alphabetical order.

It is important to note that the order of the keys is not guaranteed by the language specification and should not be relied upon in your code.
If you need to iterate over the map in a specific order, you should sort the keys explicitly before iterating over them.
*/
