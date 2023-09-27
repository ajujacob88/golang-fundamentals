package main

import "fmt"

func main() {

	fmt.Println("-----passing arrays")
	arrays()

	fmt.Println("\n-----Now Passing slices")
	slices()

	fmt.Println("\n-----Now Passing Maps")
	maps() //map element is not an addressable value,, map is addressable(means we can take the address), but inside the map the map elements are not addressable(means the address of key and value is not accessible)

	fmt.Println("\n-----Now Passing Structs")
	structs()
	//using structs is very efficient
}

func arrays() {
	nums := [...]int{1, 2, 3}
	incr(nums)
	//fmt.Printf("the address of the nums array: %p\n", &nums)
	fmt.Println("the nums array isthe same, not changed", nums)

	//while using pointers.... //but unless necessary, dont use pointers in array, instead use slices
	incrPtr(&nums)

	fmt.Println("while using pointers, the nums array changes", nums)

}

func incr(nums [3]int) {
	//fmt.Printf("the address of the nums array in the function, look its different: %p\n", &nums)
	for i := range nums {
		nums[i]++
	}
}

func incrPtr(nums *[3]int) {
	//fmt.Printf("the address of the nums array in the function, look its different: %p\n", &nums)
	for i := range nums {
		nums[i]++ //here it is actually (*nums)[i]++,, but go automatically uses *, so no need to type manually *  here..

	}
}

//------------------SLICES---

func slices() {
	nums := []int{10, 11, 12}
	incrSlices(nums)
	fmt.Println("looks the original nums slices changed", nums, "\n but if we want to add additional elements ( append) , then the original slice doesnt change")

	incrSlicesPtr(&nums) //but slices are not meant to be used with pointers, this is just for demonstration, dont use pointers in slices
	fmt.Println("but slices are not meant to be used with pointers, this is just for demonstration, dont use pointers in slices", nums)
}

func incrSlices(nums []int) { //so as we previously said, a slice already contains a pointer to the slice header, so you dont need to use a pointer to a slice to modify its elements
	for i := range nums {
		nums[i]++
	}

	nums = append(nums, 99, 100, 101)
}

// but slices are not meant to be used with pointers, this is just for demonstration, dont use pointers in slices
func incrSlicesPtr(nums *[]int) { //so as we previously said, a slice already contains a pointer to the slice header, so you dont need to use a pointer to a slice to modify its elements

	v := *nums
	for i := range v {
		v[i]++
	}

	*nums = append(*nums, 110, 120, 131)
}

//------MAPS---

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println("the original map changed ", confused, "A map value is already a pointer, so donot use pointers to map values")

}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["additional"] = 5
}

// ---------STRUCTS------
type house struct {
	name string
	room int
}

func structs() {
	myHouse := house{name: "Manith", room: 25}

	addRoom(myHouse)
	fmt.Println("look, the original struct has no change", myHouse)

	//so now pass as pointers
	addRoomPtr(&myHouse)
	fmt.Println("by passing as pointers now it changed", myHouse)
}

func addRoom(h house) {
	h.room++
}

func addRoomPtr(h *house) {
	h.room++ //no need to use indirectional operator here also, meaning no need to use (*h).room++ , go automatically does this behind the scenes
	//(*h).room++

}
