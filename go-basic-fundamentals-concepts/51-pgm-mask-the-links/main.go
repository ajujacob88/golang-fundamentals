package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(`Enter the text in argument, eg: go run main.go "aju jacob http://aju.com my name"`)
		return
	}
	var (
		text   = args[0]
		size   = len(text)
		buffer = make([]byte, 0, size)
		//buffer = []byte(text)[:0]  //another way

		in bool
	)

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*' //when you  are working with bytes, continue working with bytes..do nor convert string to byte or viceversa unless necessary
	)

	for i := 0; i < size; i++ {
		//buffer = append(buffer, text[i])
		//fmt.Println(string(buffer))

		//if i+nlink <= len(text) && text[i:i+nlink] == link {
		//if len(text[i:]) >= nlink{
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			//fmt.Println(text[i : i+nlink])
			in = true
			buffer = append(buffer, link...)
			i += nlink
		}
		c := text[i]
		switch c {
		case ' ', '\n', '\t':
			in = false //disable the masking when the next byte is white space
		}
		if in {
			c = mask
		}
		buffer = append(buffer, c)

	}
	fmt.Println(string(buffer))

	//also there is a unicode package...eg:= unicode.ToUpper()

}
