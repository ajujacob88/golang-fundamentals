package main

import "fmt"

func main() {
	name := map[string]string{
		"Aj":  "Jac",
		"Rah": "R",
		"Eld": "Koll",
		"An":  "Jac",
	}

	fmt.Println("name is", name)
	fullnam := name

	fullnam["Nik"] = "Mar"
	fullnam["Eld"] = "SKP"

	fmt.Println("Now name is also changed", name)
	fmt.Println("full name is", name)

	//difference between map and slice
	//behind the scenes, in slice a slice value is a slice header, it stores the slice header directly in itself. The pointer field in the slice header points to a backing array,, so s slice value indirectly points to a backing array
	//however in maps, a map value doesnt contain a map header in itself. Unlike a slice value, a map value is only a pointer to the map header. so a map value itself is apointer.
	//Unlike a slice value a map value is only a pointer to a map header.So a map value itself is a pointer.In turn the map header contains another pointer to the real data.So there are multiple interactions from a map value to the actual map data but the map value itself is not a structure unlike a slice.
	//So why go creators how designed the slices and maps so differently.It's because a map header is a complex and a larger data structure than a simple slice header unlike a slice.The map data structure has a lot of moving parts behind the scenes.So rather than storing it in a value it is better to store only a pointer to the map header so value assign a map value to a map variable or pass it to a function.You only pass a pointer to a map header in summary and map value is a pointer to a map header and map header tracks everything about a map Valley.So this means that passing maps to a function or assigning them to a variable is very cheap and efficient. You can pass them around as much as you want like slice values map values.Also don't carry all the keys and values with themselves.

	//so here, you cannoot create fullnam map just by copying the original map. Doing so only copies the pointer of the original map. You need a seperate map value with a new map header. SO instead of copying the map pointer, let create a new map using make function or short declaration

	//cloning
	fullname2 := make(map[string]string, len(name)) //len(name) is not necessary
	//or

	//lets clone the values manually
	for k, v := range name {
		fullname2[k] = v
	}

	fullname3 := map[string]string{}
	for k, v := range name {
		fullname3[v] = k //value of name is the key here  //but jac is repeated twice here, so takes only one automatically
	}

	fmt.Println("name", name, "\nfullname2", fullname2, "\nfullname3", fullname3)
	fmt.Println("now these 3 are seperate")

	fmt.Println()
	//Delete
	place := map[string]string{"Aj": "Mvp", "Eld": "Pbvr"}
	fmt.Println(place)
	delete(place, "Eld")
	fmt.Println(place)
	delete(place, "nothing")
	delete(place, "Nik")
	fmt.Println(place)

	//destroy a whole map / delete a whole map
	for k := range place {
		delete(place, k)
	}
	fmt.Println(place)

}
