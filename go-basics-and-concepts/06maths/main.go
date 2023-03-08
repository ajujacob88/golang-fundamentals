package main

import (
	"fmt"

	"math/rand"
	//	"crypto/rand"
)

func main() {
	//ref: https://courses.learncodeonline.in/learn/home/Go-with-Golang/section/233283/lesson/1366670?

	fmt.Println("Welcome to maths in golang")
	/*
		var mynumberone int = 2
		var mynumbertwo float64 = 5.5
		fmt.Println("The sum is: ", mynumberone+int(mynumbertwo))
	*/

	//random number   https://pkg.go.dev/math/rand#Intn

	fmt.Println(rand.Intn(5))
	fmt.Println(rand.Intn(5) + 3)

	//random from crypto

	//	fmt.Println("Random no generated from crypto")
	//	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	//	fmt.Println(myRandomNum)
}
