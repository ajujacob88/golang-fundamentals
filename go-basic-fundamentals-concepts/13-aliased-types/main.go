package main

import "fmt"

func main() {
	var (
		byteVal  byte
		uint8Val uint8
		intVal   int
	)
	uint8Val = byteVal //this works, because byte is an alias of uint8,byte is just another name for uint8..they are just different names with the same type,
	intVal = int(uint8Val)

	fmt.Println(intVal)

	var (
		runeVal  rune
		int32Val int32
	)
	runeVal = int32Val //this also works, because rune is an alias of int32,, rune is just another name for int32..they are just different names with the same type, so rune can int32 can be used together without type conversion,, again rune is used for representing unicode character
	_ = runeVal
}
