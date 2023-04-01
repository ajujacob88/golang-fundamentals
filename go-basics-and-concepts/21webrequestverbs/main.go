//https://www.youtube.com/watch?v=xh79JXJy0yY&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=28
//https://www.youtube.com/watch?v=V-sxFQ0fWlw&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=29
//program for both these videos

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welocme to web verb video")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	checkNilErr(err)
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	//content is in the byte format, so need to convert to string to read it
	//fmt.Println("content of the url is:", string(content))
	//another way of converting byte to string
	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)
	fmt.Println("byte count is: ", byteCount)
	fmt.Println("Data in string format is", responseString.String()) //this is good because we are holding always the raw data, not converting
	fmt.Println(responseString.Len())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
