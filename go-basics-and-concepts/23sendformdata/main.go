//https://www.youtube.com/watch?v=U_LjyX4iDbU&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=30
// All these concepts in these programs(not only this program) are used to work on any server given to you or any specification via the server given to you..

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("This is verb video part 2 - post Form request")

	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	//formdata, we can also send images using this

	data := url.Values{}
	data.Add("firstname", "rahul ")
	data.Add("lastname", "r")
	data.Add("age", "33")

	response, err := http.PostForm(myurl, data) //PostForm issues a POST to the specified URL, with data's keys and values URL-encoded as the request body.
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//now read that data from server
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
