package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	uniquewords := make(map[string]bool)
	for in.Scan() {
		//fmt.Println(in.Text())
		//uniquewords[in.Text()] = true //better to save all words as lower case only, so below line
		//uniquewords[strings.ToLower(in.Text())] = true  //its better to save to the map if the length of the word have atleast 3 charcters, so below step

		word := strings.ToLower(in.Text())
		if len(word) > 2 {
			uniquewords[word] = true //the words will be unique words, because the maps key can have  only unique value, if 2 same words come, then only will be accepted as key,, = ture this will makes the map a set
		}
	}

	//fmt.Println(uniquewords)

	//now we can query the words in the set

	query := "sun"
	// if _, ok := uniquewords[query]; ok {
	// 	fmt.Printf("The input contains %q.\n", query)
	// 	return
	// }
	// fmt.Printf("The input doesnot contain the word %q.\n", query)

	//but in simply we can write below, no need of the large if statement above,no need to use the ,ok idion here, instead i can directly check

	//fmt.Println("sun:", uniquewords["sun"], "aju:", uniquewords["aju"])
	if uniquewords[query] {
		fmt.Printf("Here, The input contains %q.\n", query)
		return
	}
	fmt.Printf("Here, The input doesnot contain the word %q.\n", query)

	//now gets the word from user and piping and explained in next pgm

}
