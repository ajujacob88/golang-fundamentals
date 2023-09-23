package main

//read from command line

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin) //in means input stream

	in.Scan() //default new line or /n is considered as the end of current scan

	//fmt.Println(in.Scan())
	fmt.Println("Scanned Text:", in.Text())
	fmt.Println("Scanned Bytes:", in.Bytes()) //this will be handy if you  are working with bytes only

	//now type again after printing the above 2 statements
	in.Scan()
	in.Scan()
	in.Scan()
	fmt.Println("Scanned Text 2:", in.Text()) // but this returns only the last scanned statement

	//again read all the lines from input data, for loop can also be used,
	//here in below statement, whenver i click enter it returns that data and ask me to enter more data

	for in.Scan() { //in.scan returns bool
		fmt.Println("Scanned Text 3", in.Text())
	}

}
