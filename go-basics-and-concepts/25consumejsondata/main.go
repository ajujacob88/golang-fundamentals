//Now going to consume json data,  decode json data by golang ,, decoding the json data.
//https://www.youtube.com/watch?v=a96veXdifys&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=32

package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	Price    int    `json:"courseprice"`
	Platform string
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	DecodeJson()
}

func DecodeJson() {
	//now just assume that the below json data is coming from web, actually this json data is copied from last pgm output, here purpose is to decode json data, so just entered a json data and going to decode..
	jsonDataFromWeb := []byte(`	
	{
			"coursename": "ReactJS Bootcamp",
			"courseprice": 299,
			"Platform": "learncodeonline.in",
			"tags": ["web dev","js","react"]
	}	
	`)
	//to verify the above data is in correct json format, that process is given below

	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb) //returns bool //build in function in go for chking the json data is valid or not

	if checkValid { //means checkvalid == true
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //pasiing on refernee, not copy
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//after analysing the output , notice that how smart it is that while receiving the aliases are automatically pulling up as struct names
	//so above is the common syntax and majority of cases used the  method as discussed above.

	//some cases where you just want to add data to key value pair.
	//always remember, whenever we are creating a map for online json, the key the first value is always a string, but the value can be a string/int/array/anything,, so use interface
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("\nAnother method:")
	fmt.Printf("%#v\n\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
