package main

import (
	"bufio"
	"fmt"
	"os"
)

//the same pgm can be used to read data from terminal or from file
//to read data from terminal, just run go run main.go and after than enter the text
//to read text from file, run go run main.go <filename.txt
// here eg: go run main.go < proverbs.txt
//Scan function doesnt care where it gets the input from..
//By using this symbol < i have redirected the standard input to the contents of the proverbs.txt file. This is called the command redirection

func main() {

	//os.Stdin.Close()  //i provide tit here , inorder to print the error

	in := bufio.NewScanner(os.Stdin)

	//in.Scan() //this will scan only the last line

	//to check no of lines in the file
	var lines int

	for in.Scan() {
		fmt.Println("Scanned text: ", in.Text())

		lines++

	}

	fmt.Println("No of lines in the file", lines)

	// so we have created the program that can count the no of lines from any input string

	//Again we can error check using the errr method
	if err := in.Err(); err != nil {
		fmt.Println("Error is:", err)
	}
}
