package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := `Aj✊ Ja 🙏🙏
		fdshfskAj✊ Ja 
		A🙏🙏j✊ Ja 🙏🙏`

	fmt.Println(str)

	for i := range str {
		//fmt.Println(str[i])
		fmt.Printf("%c", str[i])

	}
	fmt.Println()
	byteslic := []byte(str)
	fmt.Println(byteslic)
	for i := range byteslic {
		//fmt.Println(str[i])
		fmt.Printf("%c", byteslic[i])

	}
	fmt.Println()

	runeslic := []rune(str)
	fmt.Println(runeslic)
	fmt.Println(string(runeslic[2]))
	for i := range runeslic {
		//fmt.Println(str[i])
		fmt.Printf("%c", runeslic[i])

	}

	fmt.Println()

	//how to decode this easily
	//go offers us , for range loops as a way to decoding the runes automatically in a string

	for _, v := range str {
		fmt.Print(string(v))
	}
	fmt.Println()
	//or
	for _, v := range str {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	newstr := strings.ToLower(str[:1])
	fmt.Println(newstr)

	//in the word below ω is a greek letter, so its size is 2 bytes.. inorder to convert it to upper case only that letter , we need to pass the bytes of that charcter, so it is like given below
	//in go standard library, every package or functions deals with bytes will also have similar one with strings and vice versa, because both are almost same,,, so like strings.ToUpper, bytes.ToUpper is also present
	word := []byte("ωa🙏🙏jU")

	fmt.Println(string(word))
	r, size := utf8.DecodeRune(word)
	fmt.Println(r, size)
	fmt.Printf("%c size:%v\n", r, size)
	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Println(string(word))

	//Again, slicing a string is very efficient and cheap.. it reuses the same backing array
	//converting  string to bytes orrune is a costlier operation as it crreates new backing array.. so do it only if it is necesary..so unless necessary, do not convert them

}
