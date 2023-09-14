// a unique number generator, limit number entered via terminal
package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Enter the limit also, eg: go run main.go 25")
	}

	fmt.Println(args)
	// Seed your custom random generator with a specific seed value. //use the below instead of depreciated rand.seed
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	max, _ := strconv.Atoi(args[0])

	var uniques []int
loop:
	for len(uniques) < max {
		n := myRand.Intn(max) + 1
		fmt.Print(n, " ")
		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)

	}

	fmt.Println("\n\n Uniques:", uniques)

	//now to sort the array, use sort package fromm standard library

	sort.Ints(uniques)
	fmt.Println("sorted is:", uniques)

}
