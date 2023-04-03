//https://www.youtube.com/watch?v=h5NeKZuzUoc&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=32

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("This is verb video part 2 - post request with json data")

	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/postaju"

	//create some of the data, server accepts the data in json format or in the url encoded form.
	//in this eg, we are sending json data, but we didnt yet studied the json data, which will study after few videos.
	//right now doing a simple hack of creating a fake json data or json payload

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"myname" : "Aju Jacob",
		"Address" : "Manithottam",
		"age" : 33
	}		
	`)

	//now send the request, use http, just like get we have post in http
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// now rading that data from the server(as like in previous program verbs)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	//so we sent json data to the server and received that sended data from the server..
}
