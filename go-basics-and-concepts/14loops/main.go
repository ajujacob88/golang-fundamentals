package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in go")
	//ref: https://youtu.be/ZWBA3l818y0

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//type 1

	for i := 0; i < len(days); i++ { //we was not able to use ++i in golang
		fmt.Println(days[i])
	}

	//another type
	fmt.Println("Another type")

	for i := range days { //keyword range and this range automatically loops through over any array or slice, in this case slice,, here the range is passes as index
		fmt.Println(days[i]) //here is is getting the index only and not the values
	}

	//another type
	fmt.Println("Another type")

	for i, dayaju := range days { //here dayaju is getting individual values itself
		fmt.Printf("Index is %v and value is %v\n", i, dayaju)
	}

	//if you are not interested in index or i, you can replace it with _
	fmt.Println("if you are not interested in index or i, you can replace it with _")
	for _, dayaju := range days { //here dayaju is getting individual values itself
		fmt.Printf("value is %v\n", dayaju)
	} // this one will be commonly ised, first one returns index and then the value by using range. if you are interested in index use the first parameter by changing _, or if we need value, then use the second parameter

	//Another thing
	fmt.Println("Another type ")
	ajuValue := 3

	for ajuValue < 10 {
		fmt.Println("Value is ", ajuValue)
		ajuValue++ //++ajuvalue is not possible in go
	}

	//Another thing
	fmt.Println("Break and continue")
	ajuValue2 := 2

	for ajuValue2 < 10 {
		if ajuValue2 == 6 {
			break
		}
		fmt.Println("Now it is ", ajuValue2)
		ajuValue2++ //++ajuvalue is not possible in go
	}

	fmt.Println("Continue")
	ajuValue3 := 2

	for ajuValue3 < 10 {
		if ajuValue3 == 6 {
			ajuValue3++
			continue
		}
		fmt.Println("Now using continue it is ", ajuValue3)
		ajuValue3++ //++ajuvalue is not possible in go
	}

	//goto statement
	fmt.Println("goto statement")

	ajuValue4 := 2

	for ajuValue4 < 10 {
		if ajuValue4 == 6 {
			goto jumpaju
		}
		fmt.Println("before jumpinge it is ", ajuValue4)
		ajuValue4++
	}

jumpaju:
	fmt.Println("Jumping to aju ")
}
