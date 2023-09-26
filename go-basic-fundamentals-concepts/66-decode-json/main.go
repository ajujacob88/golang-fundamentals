// decode data from a json file
// first i read it from json and copy it to byte slice and then decode it to a users slice
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// type user struct {
// 	Name        string
// 	Permissions map[string]bool
// }

// here in struct the field name is Name, but in json the field name is username, so you need to tell it the json package, to do that we need to do the fields tags, `json:"username"`
type user struct {
	Name        string `json:"username"`
	Permissions map[string]bool
}

func main() {
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		//input = append(input, '\n')  //bufio.scanner removes the new line, thats why while printing it looks awkword, you can verify this by appending new line,, but this is unnecessary
	}

	//fmt.Println(string(input))
	//so we got the json, now its time to decode it

	var users []user
	err := json.Unmarshal(input, &users) //unmarshal decodes json to a value,, here the unmarshal function decodes the json to the users slice,,, also pointer &users is required because everything in go is pass by value, afunction cannot change a variable directly, here the users slice is a variable, the function cannot change that. There is only one way for a function to change a variable, that is by using pointers and hence users memory address is passed using &
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)

	fmt.Println()

	//or we can use also like below
	for _, user := range users {
		fmt.Println(user.Name)
	}
}
