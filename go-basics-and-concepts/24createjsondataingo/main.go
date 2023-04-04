//https://www.youtube.com/watch?v=SZ5xZ9OTeEI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=31
//in the last couple of programs, we received the json data, but we not treated it as json data, instead we treated it as string
//Now we are going to take care of the json. It is very crucial and essential part that you have to understand how to create json and how to hanfdle json when it is coming from a web request.

// in this pgm, going to work with the encoding of the json. This simply means, I have a data maybe sting or slices or array or anything or maps or anything, I wanted to convert the data into a valid json. That process is usually known as the encoding of the json. That is the goal here.
package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"` //creating aliases,we can call it as Name and inside the API if we eed to call it as coursename, we can put that b/w ``, not compulsory but best practice
	Price    int      `json:"courseprice"`
	Platform string   //aliasing not compulsory
	password string   `json:"-"`              //putting up - inside double quotes, - simply means I dont want this field to reflected whoever is consuming my APIs. this way we can remove the pasword showing there
	Tags     []string `json:"tags,omitempty"` //omitempty says, if the valie is nil or nul, then just dont throw that field at all. aslo make sure no space is allowed before omitempty, otherwise gives error
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []courses{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web dev", "js", "react"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "bcd123", []string{"full-stack", "js", "MERN stack"}},
		{"Angular Bootcamp", 299, "learncodeonline.in", "hit123", nil},
	}

	//goal is to package this data as JSON data

	finalJson, err := json.Marshal(lcoCourses)

	if err != nil { //the above process is very prone to error
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

	//but here the json data is not so readable, then password is showing and null is showing in tags with nil value,,, so for better formating we can use MarshalIndent
	fmt.Println("Now using Marshal Indent")

	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t") //here the result in indented based on \t (tab),also we can use space or anything
	if err != nil {                                             //the above process is very prone to error
		panic(err)
	}
	fmt.Printf("%s\n", finalJson2)

}

//Now going to consume json data,  decode json data by golang ,, the json data to server sent here is going to receive and decode that json data.
//https://www.youtube.com/watch?v=a96veXdifys&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=32

func DecodeJson() {

}
