package main

import "fmt"

type text struct {
	title, words string
	pages        int
}

type book struct {
	isbn  string
	title string //title is inside ttext struct also, when using the anonymous fields and if the field names conflict,the parent types takes priority.. here book is the parent type or outer type and text is the embedded type or inner type
	text         //anonymous struct embedding
}

func main() {

	ajubook := book{
		isbn:  "asdhjasdjdsj12",
		title: "THIS WILL GET PRIORITY",
		text:  text{title: "this wont get priority", words: "sample", pages: 520},
	}
	fmt.Println(ajubook)

	ajubook.title = "WHILE ANONYMOUS CALLING THIS WILL GET PRORITYYYYYYYY"

	fmt.Println("\n", ajubook)

	ajubook.text.title = "Now this will be changed, if called via name then this will change"
	fmt.Println("\n", ajubook)
}
