package main

import (
	"fmt"
	"unsafe"
)

type ajucollection []int

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	arrchange(array)
	fmt.Println("after change fn, original array is not changed:", array) //which means while passing array directly to a function, its actually the copy of that array will get into that function

	//Now consider the case of slice

	slice := []int{11, 12, 13, 14, 15, 16, 17}
	slicechange(slice)
	fmt.Println("after change fn, original slice is ALSO CHANGED:", slice) //this is because, the slice header consists of only the slice's back arrays starting address, size and capacity.., which means go makes a new copy of the slice header only and passes it to the function,But it doesn't make the copy of the backing array, so the slices still using the same backing array

	slice2 := ajucollection{21, 22, 23, 24, 25, 26, 27}
	slicechange2(slice2)
	fmt.Println("after change fn, original slice is ALSO CHANGED:", slice2)

	fmt.Printf("here arrays size: %d bytes(array size 6 * 8 bytes=48)\n", unsafe.Sizeof(array))
	fmt.Printf("slices size:%d bytes,, \nthis is actually slice header size,it will be always(3*8=24, sliceheader conists of pointer(size 8), length(size 8) and capacity(size 8), its backing arrays slice will be according to the numnber of elements\n) ", unsafe.Sizeof(slice))

}

func arrchange(array [6]int) {
	fmt.Println("original array:", array)
	array[2] = 3333
	fmt.Println("changed array:", array)
}

func slicechange(slice []int) {
	fmt.Println("original slice:", slice)
	slice[2] = 3333
	fmt.Println("changed slice:", slice)
}

func slicechange2(again ajucollection) {
	fmt.Println("original slice:", again)
	again[2] = 3333
	fmt.Println("changed slice:", again)
}
