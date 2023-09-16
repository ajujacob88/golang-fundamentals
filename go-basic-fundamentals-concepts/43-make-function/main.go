package main

import (
	"fmt"
	"strings"
)

// by using make, we can allocate a slice with a preallocated backing array.. we can optimise the code if make is used wisely, otherwise short declaration will be good.
// make initialises and returns a slice withthe given length and capacity eg: make([]int, 3) - here 3 is the size and capacity, make([]int, length, capacity) length & capacity is the length and capacity of the backing array
// we can use the make function if we know the length of the slice.. we can use it to allocate a new backing array with any length and capacity
func main() {
	//here the task is to store the slice in capital lette rto anothe slice
	tasks := []string{"jump", "run", "read"}

	//var upperTasks []string
	//but here we actually know the size of the uppertasks slice, so using var upperTasks []string will be more costlier operation, becauuse backarray will be created  while appending, since we know its length and capacity, here we can use make function

	upperTasks := make([]string, 0, len(tasks)) //here len(tasks) is 3, so this willl make a backarray with length 0 and capacity 3, using make has effect only in memory allocation, so we can view the difference in the terminal
	for _, v := range tasks {
		upperTasks = append(upperTasks, strings.ToUpper(v))
	}
	fmt.Println(tasks)
	fmt.Println(upperTasks)

	//There is abuilt in copy function to copy the element sto slice,
	//the copies functioncopies the elements based on THE LENGTH OF THE SMALLEST SLICE, also its type should be same
	dest := []int{1, 2, 3, 4, 5, 6}
	source := []int{11, 12}
	fmt.Println(dest)
	copy(dest, source)
	fmt.Println(dest)

	dest1 := []int{20, 21}
	source1 := []int{22, 23, 24, 25}
	fmt.Println(dest1)
	copy(dest1, source1)
	fmt.Println(dest1)

	dest2 := []int{30, 31, 32, 33, 34}
	fmt.Println(dest2)
	copy(dest2, []int{4, 5, 6})
	fmt.Println(dest2)

	//also the copy function return an int value which tells how many elements are copied
	dest3 := []int{}
	fmt.Println(dest3)
	n := copy(dest3, []int{4, 5, 6})
	fmt.Println("dest3 is", dest3, "n is", n)

	//but inorder to copy from a same  size of slice, then use make also
	sourceslice1 := []float64{40, 41, 42, 43, 44, 45}
	destslice1 := make([]float64, len(sourceslice1))
	n1 := copy(destslice1, sourceslice1)
	fmt.Println("source:", sourceslice1)
	fmt.Println("dest:", destslice1, "n:", n1)

	//but in realtime it is better to use append function instead of copy function in the case of slice, because its is easy and no need to use make function there
	name := []string{"aju", "an", "nik", "rah", "ras"}
	namecopy := append([]string{}, name...)
	fmt.Println("name:", name)
	fmt.Println("namecopy", namecopy)

}
