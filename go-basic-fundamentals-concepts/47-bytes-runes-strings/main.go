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

}
