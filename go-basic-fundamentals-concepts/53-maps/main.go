package main

import (
	"fmt"
	"os"
)

func main() {
	var dict map[string]string
	fmt.Printf("zero values: %#v\n", dict)
	fmt.Printf("# of keys: %d\n", len(dict))

	var complexmap map[string][]float64 //the key value cant be a slice or array, a map key should be a comparable type.. Also dont use a float type as a map key as it may lead to inaccurate results,instead int or string keys because they are more efficient

	_ = complexmap

	key := "good"
	value := dict[key]
	fmt.Printf("Key is %q and value is %#v\n", key, value) //so of no value is given default value will be "" for string, 0 for int and false for bool

	//mapValue[key] = value adds or overwrites an element
	//You can onlychange a map if it is initialized
	//initialising the map
	// dict["good"] = "best"
	// dict["happy"] = "sad"
	// dict["hi"] = "hello"
	// dict["good"] = "replacing"
	// fmt.Println(dict)

	check := map[string]string{}
	check["good"] = "best"
	check["happy"] = "sad"
	check["hi"] = "hello"
	check["good"] = "replacing"
	fmt.Println(check)

	name := map[string]string{} //just initialises the map with no data
	args := os.Args[1:]
	first := args[0]

	name["Aj"] = "Jac"
	name["Paul"] = "Pet"
	name["mistake"] = ""

	fulvalue, ok := name[first] //here ok is bool value which returns true if the key exists,, if we dont need that, then just omit it like this fulvalue := name[first]
	// if fulvalue == "" {
	// 	fmt.Printf("%q name not found\n", first)
	// 	return
	// }

	if !ok {
		fmt.Printf("%q name not foundd\n", first)
		//return
	}

	//another way
	if val4, ok := name[first]; !ok {
		fmt.Printf("%q name not foundd \n", first)
		return
	} else {
		fmt.Println(val4)
	}

	fmt.Printf("full name of %v is %v \n", first, fulvalue)

	mapdec := map[int]string{ //initialises the map with key value pair
		35: "Aj",
		31: "Nik",
		45: "Rah",
	}

	mapdec[45] = "sree"

	fmt.Println(mapdec)

	fmt.Println(mapdec[31])
	fmt.Println(mapdec[56])

	val3, ok := mapdec[31]
	fmt.Printf("value is %q and ok is %v\n", val3, ok)

	// map is designed for fast lookup, not for traversal..anyway in special cases especially in testing, if we need to traverse the map, then we can use for loop,, but remember the index order will be differnet everytime
	for i, v := range mapdec {
		fmt.Printf("key: %v and value: %q\n", i, v)
	}

	fmt.Println()
	//we can directly print the key, please note the go fmt package always prints the maps in sorted order of key automatically
	fmt.Println(mapdec)
	fmt.Printf("%#v\n", mapdec)

	//check equality using sprintf,, remember sprintf is similar to printf, but instead of printing it returns a string value
	//now i am foing to copy a map
	copied := map[int]string{31: "Nik", 35: "Aj", 45: "sree"}

	orig := fmt.Sprintf("%v", mapdec)
	new := fmt.Sprintf("%v", copied)

	if orig == new {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps NOT equal")
	}

}
