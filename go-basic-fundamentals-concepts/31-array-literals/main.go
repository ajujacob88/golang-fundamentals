package main

//literals and ELLIPSIS
import "fmt"

func main() {

	//this is the array literal,, creates and initialise  a new array with the given values
	var books = [4]string{
		"harr", "tez", "hiih", "sdasd",
	}

	fmt.Printf("%q\n", books)

	//short declaration can also be used

	quant := [4]int{
		3, 4, 5, 2,
	}
	fmt.Printf("%v\n", quant)

	booklat := [4]string{"as", "fsd"}
	fmt.Printf("%q\n", booklat)

	//ELLIPSIS - by using the three dots, go findout the lengthof the array automatically

	booklatest := [...]string{"ah", "sasad"}
	fmt.Printf("%q\n", booklatest)

	//booklatest[3]= "asd"
}
