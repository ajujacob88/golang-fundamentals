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
	//ioutil provides functions for working with files, ioutil package is depreciated, so use os package or io package
	//i am going to read the names of the file inside the folder
	//after that i am going to write the file names to a file

	fileinfo, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var names []byte //os.WriteFile accepts only byte values. A string value is actually a read only byte slice.
	//var names2 []string

	for _, file := range fileinfo {
		filename := file.Name()
		fmt.Println(filename)
		//fmt.Println(file.Type())
		//fmt.Println(file.Info())
		names = append(names, filename...) //the name string can be converted to byte values using ... here the string value is converted to a bytle slice on the fly..using ellipse operator the elements in the byte slice to the main slice
		names = append(names, "\n"...)

		//	names2 = append(names2, filename)
		//	names2 = append(names2, "\n")

	}
	fmt.Println(names)
	fmt.Printf("%s", names) //the s verb converts the given byte value to a string value automaticcally on the fly
	//fmt.Println(names2)

	err = os.WriteFile("demofiles/aju2.txt", names, 0644) //perm is the file permissions, the value can be found out using chmodcalculator or file permission calculator in web, from there i selected the file permissions and we will get the code of 0644 or anything according to out set file permissions
	if err != nil {
		fmt.Println("write error: ", err)
		return
	}

}
