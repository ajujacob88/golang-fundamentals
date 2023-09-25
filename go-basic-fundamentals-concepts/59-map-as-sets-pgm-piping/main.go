package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// pgm same as 58th one, but added more things like piping

// we can even check a website that contains the unique word,, pipe the curl command to our pgm,
// for that run like this :  curl -s  https://www.rfc-editor.org/rfc/rfc20.txt | go run main.go outside  //give any web url , and rplace outside with the word to query

// lastly added regular expressions (regexp) to eliminate the quotes, numbers, symbols etc from text. Regular expressions finds patterns in a text
func main() {
	//creating a regexp patterns that matches any characters but letters.
	rx := regexp.MustCompile(`[^a-z]+`) //here ^ means do not match, a-z means a to z characters, + issame pattern 1 or more times..so this pattern means find any character except letters 1 or more times

	//user input
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter the query as input also, eg: go run main.go sun <shakespeare.txt")
		return
	}
	query := args[0]

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	uniquewords := make(map[string]bool)
	for in.Scan() {

		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "") //here replacing the characters found via regexp with an empty string.. this will remove anything except letters

		if len(word) > 2 {
			uniquewords[word] = true //the words will be unique words, because the maps key can have  only unique value, if 2 same words come, then only will be accepted as key,, = ture this will makes the map a set
		}
	}

	//query := "sun"

	if uniquewords[query] {
		fmt.Printf("Here, The input contains %q.\n", query)
		return
	}
	fmt.Printf("Here, The input doesnot contain the word %q.\n", query)

}
