package main

import (
	"fmt"
)

func main() {

	/* ONE METHOD
	fmt.Println("Enter the two numbers")

	numberReader := bufio.NewReader(os.Stdin)
	input1, _ := numberReader.ReadString('\n')
	num1, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64) //this 10 is for the value of base entered which is  decimal here, if it is binary give 2, and if the string entered is hex give 16, then if we input F, it will show 15
	//num1, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)

	input2, _ := numberReader.ReadString('\n')
	num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	sum := float64(num1) + num2
	fmt.Println("The sum of ", num1, " and ", num2, " is ", sum)
	*/
	//Another Method

	var num1 int
	var num2 float64
	fmt.Println("Enter the integer no:")
	fmt.Scan(&num1)

	fmt.Println("Enter the float no:")
	fmt.Scan(&num2)

	sum := float64(num1) + num2
	fmt.Println("The sum of ", num1, " and ", num2, " is ", sum)

}
