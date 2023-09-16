package main

import (
	"fmt"
	"strconv"
	"strings"
)

// suppose we are getting the spending data from a file in sting form., in the data, each line represents each day, find he total spending on each day
func main() {
	data := ` 50 80
		100 800 450 200
		5000 100
		80 350 9000 250 550
		2000 100`

	spendings := fetch(data)

	fmt.Println(spendings)

	//now take the sum of each day
	for i, daily := range spendings {
		var total int
		for _, v := range daily {
			total += v
		}
		fmt.Printf("Day: %v, Total spends: %v\n", i, total)
	}

}

func fetch(data string) [][]int {

	//the data is in the form of string, so i need to convert it into int and then store in the slice as each day
	lines := strings.Split(data, "\n") //split the string in each line
	//fmt.Println("lines:", lines)
	spendings := make([][]int, len(lines))
	//fmt.Println("spendings init,", spendings)

	for i, lineData := range lines {
		fields := strings.Fields(lineData)
		//fmt.Println("fields", fields)
		spendings[i] = make([]int, len(fields))

		for j, v := range fields {
			datainInt, _ := strconv.Atoi(v)
			spendings[i][j] = datainInt
		}

	}
	return spendings
}
