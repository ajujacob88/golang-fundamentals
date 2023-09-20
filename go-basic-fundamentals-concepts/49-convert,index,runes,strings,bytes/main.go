package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

// string is a byteslice
// string literals ar automatically encoded in utf-8
func main() {
	str := "hii ✊ jac 🙏"
	fmt.Println(str)

	fmt.Println(str[0])

	for _, v := range str {
		fmt.Println(v)
	}

	//str[0] = 'f' //not possible, since strings are read only byte slices
	//inorder to do that,conver the string to byte slice and then change instead

	byteslice := []byte(str) //here, creating a new []byteslice and copies the bytes of the string to new slice's backing array, so they dont share the same backing array. So changing the byte slice doesn't change the original slice
	byteslice[0] = 'H'
	byteslice[1] = 'E'
	byteslice[2] = 'Y'

	fmt.Println("str:", str, "\nbyteslice:", string(byteslice), "\njust:", byteslice)
	// in summary here created a new []byteslice and copies the bytes of the string to new slice's backing array, so they dont share the same backing array. So changing the byte slice doesn't change the original slice

	str = string(byteslice) //creates a string and copiesthe []bytes to thenew strings backing array
	fmt.Println(str)

	//so in summary, a string is an immutable byte slice

	name := "Aj✊ Ja 🙏🙏"
	namebyte := []byte(name)

	fmt.Printf("length of name is %d bytes\n", len(name)) //len() function counts the bytes, not the runes, inorder to count the actual runes, use rune count in string function
	fmt.Printf("length of namebyte is%d bytes\n", len(namebyte))

	fmt.Printf("no of runes in name is  %d \n", utf8.RuneCountInString(name))
	fmt.Printf("no of runes in namebyte is  %d\n", utf8.RuneCount(namebyte)) //these mirror functions helps you to preventconverting a []byte to a string or viceversa because conversion is costly. There are lot of mirror functions like these because string and []byte  are interchangeable values
	//so here inthe above string there are 9 runes/characters and 17 bytes,, the bytes are 17 because, the non english characters takes more bytes, here the 2 emojis took 2*4 = 8 bytes and 1 emoji took 3 bytes  and the english characters took 6*1 = 6 bytes,, the bytes count for each rune is hsown below

	for i := range name {
		fmt.Printf("strin[%2d] = %-2x\n", i, name[i]) //x is hexadecimal
	}

	fmt.Println("\n now the utf8 encode version of the rune here,,,")
	for i, r := range name {
		fmt.Printf("strin[%2d] = % -12x = %q\n", i, string(r), r) //remember when we convert an integer number, u encodes it to utf 8 automatically
	}

	fmt.Println("1st byte/rune: ", string(name[0]))
	fmt.Println("2nd byte/rune: ", string(name[1]))
	fmt.Println("3rd byte: ", string(name[2]))            //the third byte ✊ is not printed correctly using index, because it is 3 byte length, so name[2],name[3] and name[4] togetherly stored these bytes,, but we can print it by using slicing
	fmt.Println("3rd rune is: ", string(name[2:5]))       //by slicing, the ✊  can be printed correctly
	fmt.Println("last rune:", string(name[len(name)-4:])) //🙏 is 4 bytes long

	// so in summary, each rune in a string is having multiple number of bytes varies from 1 to 4 bytes. so we can't easilt index the strings containing multiple type runes like english and hindi together or emojis etc

	//in other scripting languages you can index utf8strings easily, but in go it is not allowed to do so by default.. but there is an easier way to indexing, ie by converting string to rune, but its not efficient..

	namerune := []rune(name)
	fmt.Println("1st rune: ", string(namerune[0]))
	fmt.Println("2nd rune: ", string(namerune[1]))
	fmt.Println("3rd rune: ", string(namerune[2]))
	fmt.Println("some in between runes:", string(name[1:13]))

	//but the problem is converting string to rune is inefficient, because rune is an alias of in32, hence take more space,,, so we dont use runes all the time, instead we commonly uses strings which is a byteslice
	//actually string is a byte slice, so it tooks only 1 bytes for 1 byte runes, but rune is int32, so it will take 4 bytes for each runes,
	//so if we store "aju" as a string, it will took only 3*8 = 24bytes memory,, but if "aju" is stores as a rune, it will take 3*32 = 96bytes memory..this is because rune is an alias of int32, so its actually int32
	//llly if we store "aju🙏" as a string, it will took only 3*8 + 4*8 = 56 bytes memory,, but if "aju🙏" is stores as a rune, it will take 4*32 = 128bytes memory..this is because rune is an alias of int32, so its actually int32

	fmt.Println("length:", len(namerune)) //length will be 9 because here the length of rune means no of runes/characters, so here namerune is "Aj✊ Ja 🙏🙏" which has 9 runes/characters

	fmt.Println(unsafe.Sizeof(namerune[0]))
	fmt.Println(unsafe.Sizeof(namerune[2]))
	fmt.Println(unsafe.Sizeof(name[0]))
	fmt.Println(unsafe.Sizeof(name[2]))
	fmt.Println(unsafe.Sizeof(namebyte[0]))
	fmt.Println(unsafe.Sizeof(namebyte[0]))

	// in summary, string is a byteslice
	// string literals ar automatically encoded in utf-8,, only assume that the string literals are encoded in utf8
	//we can store anything in a string value, beause it only contains bytes, its a byte slice,, so on a string value we can even store an image, a music or sound and so on,, so every string value is not a urf 8 string value, only the string literals are utf8 encoded by defualt.
	//single byte runes are also called ASCII runes
	//strings are immutable, means we cannot change them, only we can create new values
}
