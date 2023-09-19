package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go run main.go 128512 128591       -- for emogi code points
//go run main.go 64 129000
// search the different utf8 code points and put it along go run main.go

func main() {
	// start, stop := 'A', 'Z' //rune literals.. actually they are numeric unicode code points

	// fmt.Println(start, stop) //65 is the code point of A, 90 is the code point of Z....(like ASCII values, in utf8 it is called as code points)
	// fmt.Println()

	var start, stop int
	args := os.Args[1:]
	if len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'z' //put this as default, if no input is given via console,,, we can assign rune literals (like 'A', 'B' etc) to int or any numeric type because rune literals are typeless,, each gets converted to int here
	}

	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n", "literal", "decimal", "hexadecimal", "utf8 encoded code point(in hexadecimal)", strings.Repeat("-", 50)) //adding heading for the character set table

	//now printing all the characters between start and stop points

	for i := start; i <= stop; i++ {
		//fmt.Print(i, " , ")
		//printing runes using different printf verbs
		//fmt.Printf("%c => %[1]d\n", i) //this is same as fmt.Printf("%c => %d\n", i, i)
		fmt.Printf("%-10c  %-10[1]d %-10[1]x % -12x\n", i, string(i)) //string(i) encodes the given code point to utf-8 encoded string value automatically by go. space -32 that space is to seperate single bytes
	}

	//var char byte = '✊' //not possible, because it is 3 byte long, so it overflows the byte variable
	var char1 rune = '✊'
	var char2 int = '✊'
	char3 := '✊' //default type of a rune literal is rune type
	char4 := '🙏'
	fmt.Println(char1, char2, char3, char4)

	//so in summary, usiong utf 8 we can reprent code point between 1 bytes and 4 bytes,, so i can store these code byte in a rune variable.
	//you can represent any unicode code point using the rune type because it can store 4 bytes of data(since rune is int 32),, so if you want to represent any unicode code point you can use the rune type and rune literal because it can store upto 4 bytes
	//https://blog.hubspot.com/website/what-is-utf-8    -  read this also
}
