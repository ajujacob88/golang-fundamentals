package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")
	//ref: https://youtu.be/_0R6H1m9o78

	languages := make(map[string]string)
	//adding values
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)

	fmt.Println("JS shorts for: ", languages["JS"])

	//deleting values
	delete(languages, "RB")
	fmt.Println("List of all languages after deletion: ", languages)

	// loops are interesting in golang //loops in maps

	for keyaju, valueaju := range languages {
		fmt.Printf("For key %v, value is %v \n", keyaju, valueaju) //placeholder v is for value
	}

	// we can use commaok syntax here also..., if we dont need key or the first value, then we can put _...., := walrus operator handles all these commaok syntax,, we can ignore anything at all using _, in golang

	for _, value := range languages {
		fmt.Printf("The value is %v \n", value)
	}

	// MAP IS ALSO MUTABLE DATATYPE
	var map1 map[string]int = map[string]int{"hello": 2, "hi": 3}
	map2 := map1
	map2["aju"] = 35
	fmt.Println(map1)
	fmt.Println(map2)
}
