package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb" //this is the basics of url format, 3000 or something likethat is the port no when hosted in cloud, & is used instead of comma in url

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing the url - simply means you want to extract some information from the url
	result, _ := url.Parse(myurl) //url library package

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) //here all the queries in javascript it is called as params, parameters

	queryParams := result.Query()
	fmt.Printf("The type of query params are: %T\n", queryParams) //url.Values are basically key value pairs, so we can do anything like we do in maps/dictionary since its key value pairs

	fmt.Println(queryParams["coursename"]) //in output we get the value as reactjs for the key coursename
	fmt.Println(queryParams["paymentid"])
	fmt.Println(queryParams["aju"]) //no key/value present in the name of aju

	//since its key value pair, we can use for loop and range

	for key, val := range queryParams {
		// the order is not guaranteed, since its key valu pairs(maps)
		fmt.Println("Key is:", key)
		fmt.Println("param is:", val)

	}

	//Now doing the reverse thing
	//Now construct a url if we have the information, also we dont need any libraries / packages for this
	partsOfUrl := &url.URL{ //dont forget, always pass a refernce(&) not copy of it, dont forget the ambersent
		Scheme:  "https", //this is format, no ://
		Host:    "ajujacob.com",
		Path:    "/tutorialcss",       //optional
		RawPath: "user=aju&course=go", //for rawquery
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
