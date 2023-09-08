package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	s = "how are you?" //string literal
	s = `how are you?` //raw string literal(since the quotes is backquotes ``)- both  the string and raw string literal are of type string only, raw string literal count the space as space and no need of escape sequence

	fmt.Println(s)

	s = "<html> \n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	//but this string literal is hard to read and maintain. but if raw string literal is used instead, it will be easily readable and mainainable

	s = `
<htmll>
	<body>
		"Hello"
	</body>
</html>`

	fmt.Println(s) //which means go deosnt process whats inside the string literal, so we can put next line, no need of escape characters, we can type anything

	fmt.Println("c:\\windows\\Program Files\\Aju")
	fmt.Println(`c:\windows\Program Files\\Aju`)

	// string concatenation

	first, last := "Aju", "Jacob"
	full := first + last
	fmt.Println(full)
	full = first + " " + last
	fmt.Println(full)
	fmt.Println(full + " " + "manithottam")

	full += " Mvpa"
	fmt.Println(full)

	eq := strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)

	fmt.Println(`anu` + `jacob`)

	//string length
	name := "Aju"
	fmt.Println(len(name))

}
