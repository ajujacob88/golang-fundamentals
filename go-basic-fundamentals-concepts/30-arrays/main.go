package main

import "fmt"

const (
	winter = 3
	summer = 1
	yearly = winter + summer
)

func main() {
	var books [4]string

	fmt.Println(books)

	fmt.Printf("Array Type is: %T\n", books)
	fmt.Printf("Array is: %q\n", books)

	fmt.Printf("Array is: %#v\n", books)

	var arr [yearly]int
	fmt.Printf("%#v\n", arr)

	books[0] = "harry potter"
	books[1] = "chet"
	books[2] = "hrych"
	books[3] += books[0] + " 2nd edition"

	fmt.Println(books)
	fmt.Printf("%q\n", books)
	fmt.Printf("%#v\n", books)

	//arr[5]= "test"
	//i := 5
	//arr[i] = "go compiler can only catch array indexing prolems at the compile time, here i = 5 which is greater than the size of array"

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	// wBooks[0] = arr[0]
	// wBooks[1] = arr[1]
	// wBooks[2] = arr[2]
	sBooks[0] = books[3]

	for i := range wBooks {
		wBooks[i] = books[i]
	}

	fmt.Printf("winter books: %#v\n", wBooks)
	fmt.Printf("summer books: %#v\n", sBooks)

	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("published books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %v\n", books[i])
		}
	}
}
