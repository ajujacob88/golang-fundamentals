package main

import "fmt"

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

	rune := []rune(str)
	fmt.Println(rune)
	fmt.Println(string(rune[2]))
	for i := range rune {
		//fmt.Println(str[i])
		fmt.Printf("%c", rune[i])

	}

	fmt.Println()

}
