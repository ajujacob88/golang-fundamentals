package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	change(array)
	fmt.Println("after change fn, original array becomes:", array)

}

func change(array [6]int) {
	fmt.Println("original array:", array)
	array[2] = 3333
	fmt.Println("changed array:", array)
}
