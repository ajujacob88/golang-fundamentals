package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//actually english letter is of 1 byte, for other languages, each characters can be 1 to  4 bytes long., len() gives the total byte count only, but for english its ok, since each letter are 1 byte long, but for other languages utf8 is reliable, so always utf8 is reliable to fing the no of characters of a string
	englishName := "aju jacob.."
	fmt.Println(len(englishName))

	malyalamName := "അജു"
	fmt.Println(len(malyalamName))

	hindi := "अजु"
	fmt.Println(len(hindi))

	fmt.Println(utf8.RuneCountInString(englishName))
	fmt.Println(utf8.RuneCountInString(malyalamName))
	fmt.Println(utf8.RuneCountInString(hindi))

}
