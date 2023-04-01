//https://www.youtube.com/watch?v=Mdg3tlGUXrE&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=25

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("weelcome to files in go")
	content1 := "This need to go in a text file"

	file1, err := os.Create("./mygolearningfile.txt") //to create file in the same folder

	// if err != nil {
	// 	panic(err) //panic will just shut down the execution of the pgm and will show you what the error you are facing.
	// }
	checkNilErr(err)

	length, err := io.WriteString(file1, content1)
	checkNilErr(err)
	fmt.Println("Length is: ", length)

	//file1.Close()
	defer file1.Close() //the recomended way is to use defer, because after this line yu may do some changes to the files, then if defer is not used, then it will gets closed,, here i written in last, so no need to use defer, but recommend to use deefer.

	readFile("./mygolearningfile.txt")
}

// function for reading the file
// for creation os package is used, for reading ioutil packageis there
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) //while reading the file, the data will be in binary format only..
	checkNilErr(err)
	fmt.Println("Text Data inside the file is \n", databyte)
	//so convert binary datafile to string format
	fmt.Println("Text Data inside the file is \n", string(databyte)) //this is just oeasy method, another method the tougher one is also there, which will do lateron

}

//common pattern we are using in programing for error checking, creating functions...

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
