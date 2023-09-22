package main

import (
	"fmt"
)

func main() {

	//check equality using sprintf,, remember sprintf is similar to printf, but instead of printing it returns a string value
	//now i am foing to copy a map

	namemap := map[int]string{31: "Nik", 35: "Aj", 45: "sree"}

	copied := map[int]string{31: "Nik", 35: "Aj", 45: "sree"}

	orig := fmt.Sprintf("%v", namemap)
	new := fmt.Sprintf("%v", copied)

	if orig == new {
		fmt.Println("Maps are equall")
	} else {
		fmt.Println("Maps NOT equal")
	}

	namemap2 := map[string]string{"ak": "Nik", "sd": "Aj", "sad": "sree"}

	copied2 := map[string]string{"ak": "Nik", "sd": "Aj", "saaaaad": "sree"}

	orig2 := fmt.Sprintf("%v", namemap2)
	new2 := fmt.Sprintf("%v", copied2)

	if orig2 == new2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps NOT equal")
	}

	//we cant compare namemap2 == copied2 because maps,slices etc are not comparable,, but we can compare it to nil value if namemap == nil{}
	// if namemap == copied {
	// 	fmt.Println(" maps, slice etc are not comparable type")
	// }
	if namemap == nil {
		fmt.Println("can be compared only to nil")
	}

}
