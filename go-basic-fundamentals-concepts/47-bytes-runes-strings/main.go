package main

//https://blog.hubspot.com/website/what-is-utf-8    -  read this also
import "fmt"

func main() {
	// a string is just a series of bytes
	// string and []byte are interchangaeabily convertable, hence a byte slice can be encoded to a string value

	fmt.Println(string([]byte{104, 101, 121, 65})) //converting byte value to a string,,'h' is the rune literal, 104 is the unicode points,,, in go instead of sayingunicode pints, we simply says runes,, Runes can represent a unicode code point or a simple character
	fmt.Println([]byte("heyA"))                    //converting string value to []byte

	//instead of a number, we can also represent the character by rune literals 'h' 'e' 'y'
	runeliteral := 'h'

	fmt.Println(runeliteral)

	start, stop := 'A', 'Z' //rune literals.. actually they are numeric unicode code points

	fmt.Println(start, stop) //65 is the code point of A, 90 is the code point of Z....(like ASCII values, in utf8 it is called as code points)

	var num int
	num = 'A' // we can assign rune literals (like 'A', 'B' etc) to int or any numeric type because rune literals are typeless,, each gets converted to int here. They can be assigned to a variable with any numeric type
	fmt.Println(num)
}
