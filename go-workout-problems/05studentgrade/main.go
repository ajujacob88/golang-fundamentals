package main

import "fmt"

func main() {
	fmt.Println("Enter your mark: ")
	var mark float64
	fmt.Scan(&mark)

	if mark >= 90 && mark <= 100 {
		fmt.Println("Grade: A")
	} else if mark >= 80 && mark <= 89 {
		fmt.Println("Grade: B")
	} else if mark >= 70 && mark <= 79 {
		fmt.Println("Grade: C")
	} else if mark >= 60 && mark <= 69 {
		fmt.Println("Grade: D")
	} else if mark >= 50 && mark <= 59 {
		fmt.Println("Grade: E")
	} else if mark >= 0 && mark <= 49 {
		fmt.Println("Grade: F")
	} else {
		fmt.Println("Invalid mark entered")
	}

}
