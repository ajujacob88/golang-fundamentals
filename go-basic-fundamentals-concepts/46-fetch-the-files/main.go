package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a directory, eg:= go run main.go ./demofiles or go run main.go .")
		return
	}
	//ioutil provides functions for working with files
	//ioutil.ReadDir is depreciated
	fileinfo, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range fileinfo {
		fmt.Println(file.Name())
		//fmt.Println(file.Type())
		//fmt.Println(file.Info())

	}

}
