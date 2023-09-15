package main

import (
	"fmt"
	"sort"
)

func main() {
	//Array or slice to be sliced
	// newslice := sliceable [start:stop]
	// sliceable [starting index  (slice it starting from this index):stop position(slice it up upto this element position not its index)]

	msg := []string{"h", "e", "l", "l", "o"}

	slicedmsg := msg[1:4] //easy way to remeber as said earlier, consider start as index 0,1,2.3... and stop as position 1,2,3,4....
	fmt.Println(slicedmsg)

	fmt.Println(msg[0:1])
	fmt.Println(msg[0:3])

	fmt.Println(msg[0:5]) // this is same as original, which is equal to msg[:], means, if starting from 0th index, we can omit 0, if ending upto last position, we can omit that position as describe this and below eg

	fmt.Println(msg[:])  //this is equal to msg[0:5]
	fmt.Println(msg[:3]) //this is equal to msg[0:3]
	fmt.Println(msg[1:]) //this is equal to msg[1:5]
	fmt.Println(msg[:1]) //this is equal to msg[0:1]

	fmt.Println(msg[1:2])

	fmt.Println("original msg: ", msg)
	msg = msg[1:2]
	fmt.Println("saved msg: ", msg)

	//Again
	name := []string{"aju", "anu", "nik", "sus", "jac"}
	fmt.Println(name)

	name = append(name[1:2], "manu")
	fmt.Println(name)

	place := []string{"muv", "ekm", "thr", "trv", "pkd", "ktm"}

	//now print the last 3 places
	l := len(place)
	fmt.Println(place[l-3:])

	back := []int{1, 2, 3, 4, 5, 6}
	slicing := back[1:4]
	fmt.Println(back)
	fmt.Println(slicing)

	back[1] = 500
	back[2] = 600
	fmt.Println(back)
	fmt.Println(slicing) //slicing will also change, this is because both the back and slicing share the same backing array

	back = []int{11, 12, 13, 14, 15, 16}
	fmt.Println(back)
	fmt.Println(slicing)

	//convert array to a new slice
	array := [5]int{20, 21, 22, 23, 24}
	fmt.Printf("%#v\n", array)
	toslice := array[:]
	fmt.Printf("%#v\n", toslice)

	//part of array to slice
	// slicing an array

	slicingarray := array[1:4]
	fmt.Printf("%#v\n", slicingarray)

	slicingarray[1] = 9999
	fmt.Printf("%#v\n", slicingarray)
	fmt.Printf("%#v\n", array) //original array will also change, because the original array is the backing array of that slicing array
	fmt.Printf("%#v\n", toslice)

	//slice and array both act like same...
	//so here both the grades and front share the same backing array

	//grades := [...]float64{40, 20, 30, 70, 80, 90}
	grades := []float64{40, 20, 30, 700, 800, 900}
	front := grades[:3]
	fmt.Println(front)
	fmt.Println(grades)

	sort.Float64s(front)

	fmt.Println(front)
	fmt.Println(grades)

	//how we can remove the connection, we can remove it by doing like below
	grades2 := []float64{40, 20, 30, 700, 800, 900}
	var front2 []float64
	front2 = append(front2, grades2[:3]...) //here the grades2 and front2 have seperate backing array, soits not connected

	fmt.Println("front2:")
	fmt.Println(grades2)
	fmt.Println(front2)

	sort.Float64s(front2)

	fmt.Println(front2)
	fmt.Println(grades2)

	//also we can use short declarationconst like below

	front3 := append([]float64(nil), grades2...)
	fmt.Println(front3)

	arr := [10]int{1, 2, 3}
	fmt.Println(len(front))
	fmt.Println(cap(front)) //capacity
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	trimmed1 := slice[:5]
	trimmed2 := slice[5:]
	fmt.Println(trimmed1)
	fmt.Println(trimmed2)
	fmt.Println(len(trimmed1))
	fmt.Println(cap(trimmed1))
	fmt.Println(len(trimmed2))
	fmt.Println(cap(trimmed2))

	trimmed3 := slice[2:5]
	fmt.Println(trimmed3)
	fmt.Println(len(trimmed3))
	fmt.Println(cap(trimmed3))

	fmt.Println(slice[2:5])
	fmt.Println(slice)
	fmt.Println(slice[9:9])
	fmt.Println(slice[:9])

}
