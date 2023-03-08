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
