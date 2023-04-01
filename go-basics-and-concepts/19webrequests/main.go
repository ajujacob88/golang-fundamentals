package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("This is Web Request Demo Program")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type %T\n", response) //while printing the response type it is clear that *http.Response, pointer is utilised, thus  a gurarantee is that you are not getting receiving any copy of this response. you are actually getting a response that you can manipulate further.
	defer response.Body.Close()                      //Its the programmers or callers responsibility is to close the connection. According to official documentation ReadResponse or Response.Write never  closes a connection.

	//fmt.Println("response is", response)
	//majority of the reading is done by ioutils package

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes) //here i didn,t checked the status code, but we can check status code 404 or something like by placing a condition before this...
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
