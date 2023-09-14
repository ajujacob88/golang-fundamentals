package main

import "fmt"

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

}
