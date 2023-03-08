// this is just to prepare for the review...

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var inp1 int = 2
// 	inp2 := 5
// 	out := sum(inp1, inp2)
// 	fmt.Println(out)
// }
// func sum(num1 int, num2 int) float64 {
// 	result := float64(num2) / float64(num1)
// 	return result
// }

//concurrency and go routines
// https://www.youtube.com/watch?v=ZDNtUgWRDdk
// https://www.youtube.com/watch?v=V-0ifUKCkBI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=54
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	go number("a")
	fmt.Print("hi")
	number("b")
	go number("c")
	number("d")
	fmt.Print("aju")
	number("c")
	number("d")
}
func number(str string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(str)
	}
}
*/

// Interface basics : https://www.youtube.com/watch?v=lbW-KVdIXaY

package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods

func (sq square) area() float64 {
	return sq.length * sq.length
}

func (sqr square) circum() float64 {
	return sqr.length * 4
}

func (sqr square) cube() float64 {
	return sqr.length * sqr.length * sqr.length
}

// circle methods
func (ci circle) area() float64 {
	return ci.radius * ci.radius * math.Pi
}

func (ci circle) circum() float64 {
	return 2 * math.Pi * ci.radius
}

func toPrintTheShape(shape square) {
	fmt.Printf("Area of %T is %f", shape, shape.area())
	fmt.Printf("\nCircumference of %T is %f", shape, shape.circum())
}

func toPrintTheShapeCir(shape circle) {
	fmt.Printf("\nArea of %T is %f", shape, shape.area())
	fmt.Printf("\nCircumference of %T is %f", shape, shape.circum())
}

// so an interface basically group types together based on their methods.
type allShape interface {
	area() float64
	circum() float64
	//cube() float64  //we cant use this here, because this method is not present in circle., inorder to use this define this in circle also.
}

func toPrintShapeofall(shape1 allShape) {
	fmt.Printf("\nArea of %T is %f", shape1, shape1.area())
	fmt.Printf("\nCircumference of %T is %f\n", shape1, shape1.circum())
	fmt.Println(shape1)
}

func main() {
	aju := square{20}

	ar := aju.area()
	fmt.Println(ar)

	toPrintTheShape(aju)
	rahul := circle{10}
	toPrintTheShapeCir(rahul)

	myShapes := []allShape{

		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 2.3},
		square{length: 5},
	}

	for _, v := range myShapes {
		toPrintShapeofall(v)
	}
}

//to find dns records programmatically in golang

/*

package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupIP("facebook.com")
	//fmt.Println(iprecords)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

}
*/

// revision
/*
package main

import (
	"fmt"
)

func main() {
	// fmt.Println(time.Now())
	// timeis := time.Now()
	// fmt.Println(timeis.Format("01-02-2006 15:05:04 Monday"))
	// fmt.Println(time.Date(1988, time.June, 4, 02, 30, 20, 0, time.UTC))

	var myptr *int
	var num int = 23
	myptr = &num
	fmt.Println("The memory address of num is", myptr, "and the value of num is,", *myptr)
	num2 := *myptr + 8
	fmt.Println(num2)
	*myptr = *myptr + 20
	fmt.Println(*myptr)
	fmt.Println(num)

	var slice1 []int

	slice1 = append(slice1, 3)
	fmt.Println(slice1)

	slice2 := make([]int, 4)
	slice2[0] = 5
	slice2[1] = 3
	slice2[2] = 6
	slice2[3] = 5
	fmt.Println(slice2)

	map1 := make(map[string]int)
	map1["aju"] = 23
	map1["rahul"] = 52
	fmt.Println(map1)

	aju := struct1{23, "aju"}
	fmt.Println(aju)

	for index, val := range slice2 {
		fmt.Print(index)
		fmt.Println(val)
	}

}

type struct1 struct {
	age  int
	name string
}
*/
