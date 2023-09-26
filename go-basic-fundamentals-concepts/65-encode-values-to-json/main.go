// json is a structured data exchange format.. In go encoding is also called as marshaling and decode is also called unmarshal
// finally to save json to a file: eg:=  go run main.go >savedusers.json   //means save the output of the pgm to this file
package main

import (
	"encoding/json"
	"fmt"
)

// type user struct {
// 	name        string
// 	password    string
// 	permissions map[string]bool
// }

//typing the map types like above is cumbersome, so we can declare a new map type and strucutre according to that

// type permissions map[string]bool

// type user struct {
// 	name     string
// 	password string
// 	permissions
// }

//json package encode only the exported fields(which means json package in go encodes the fiels which are starting with first letter capital), so rewritten like below

// type permissions map[string]bool

// type user struct {
// 	Name        string
// 	Password    string
// 	Permissions permissions
// }

// now adding json tags.. this json field tags are mostly used for controlling the encoding and deoding behaviour
type permissions map[string]bool

// here json: is the jey tag,, only the json package will read this tag,, json packageonly reads a field tag, if it starts with: json: ,,, other than json: there are various other field tags are there as well for thier respective packages
type user struct {
	Name     string `json:"username"` //here while encoding , i need to change the json encode field name to username, so i can use this field tag
	Password string `json:"-"`        //for seurity reasons, we dont need to encode this exported password field, then we can use json tak like this `json:"-`, then this field wont be encoded to json
	//	Permissions permissions `json:"perms,omitempty"` //here sometimes, the permissions data is nil value, so need to encode the nul value, so we can use omitempty field tag,,,here field tag use 2 things, first it changes the encoded json field name to perms, and then it omit the null data from encoding
	Permissions permissions `json:",omitempty"` //i dont want to change the name of this field while encoding, but i need to only omit the empty values from encoding, then use like this, dont foreget to put , here
}

func main() {

	// inac := user{
	// 	name:        "Inac",
	// 	password:    "123",
	// 	permissions: nil,
	// }
	// gadly := user{
	// 	name:        "Gadly",
	// 	password:    "456",
	// 	permissions: map[string]bool{"admin": true, "superadmin": false},
	// }
	// eagle := user{
	// 	name:        "Eagle",
	// 	password:    "789",
	// 	permissions: map[string]bool{"write": true},
	// }

	// allusers := []user{inac, gadly, eagle}

	//the above declaration can be simply written like below

	// users := []user{
	// 	{
	// 		name:        "Inac",
	// 		password:    "123",
	// 		permissions: nil,
	// 	},
	// 	{
	// 		name:        "Gadly",
	// 		password:    "456",
	// 		permissions: map[string]bool{"admin": true, "superadmin": false},
	// 	},
	// 	{
	// 		name:        "Eagle",
	// 		password:    "789",
	// 		permissions: map[string]bool{"write": true},
	// 	},
	// }

	//typing the map types like above is cumbersome, so we can declare a new map type and strucutre according to that
	users := []user{
		{
			Name:        "Inac",
			Password:    "123",
			Permissions: nil,
		},
		{
			Name:        "Gadly",
			Password:    "456",
			Permissions: permissions{"admin": true, "superadmin": false},
		},
		{
			Name:        "Eagle",
			Password:    "789",
			Permissions: permissions{"write": true},
		},
	}

	fmt.Println(users)

	//Now its time to encode the users slice to json
	jsonencodeddata, err := json.Marshal(users) //json is a string format  which is also a byte slice, so this returns []byte slice
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\njson encoded data is:\n", string(jsonencodeddata))

	//but here the output doesnt looks good for readability, so lets format using marshal indent function
	jsonout, err := json.MarshalIndent(users, "", "\t") //this is just for formating, here \t means a tab space
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\njson  encoded data in indented way is:\n", string(jsonout))
}
