package main

import "fmt"

func main() {
	unnamedarr := [4]int{1, 2, 3, 4} //here the type is [4]int ewhich is the unnamed type and its underlying type is the same itself which is [4]int itself
	fmt.Println(unnamedarr)

	//named array
	type bookcase [4]int //here the type is bookcase and underlying type is [4]int

	namedarr := bookcase{2, 5, 6, 8}
	fmt.Println(namedarr)

	fmt.Println(unnamedarr == namedarr)

	type cabinet [4]int
	cab := cabinet{2, 5, 6, 8}

	fmt.Println(namedarr == bookcase(cab))

	type (
		integer int
		age     [3]int
		year    [3]integer
	)

	_ = [3]integer{} == year{}

	//_ = [3]integer{} == age{}
}
