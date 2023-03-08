package main

import "fmt"

func main() {
	var mark float64
	fmt.Println("Enter the mark of the subject")
	fmt.Scan(&mark)
	if mark >= 50 && mark <= 100 {
		fmt.Println("Passed")
	} else if mark >= 0 && mark < 50 {
		fmt.Println("Failed")
	} else {
		fmt.Println("Invalid mark entered")
	}

}
