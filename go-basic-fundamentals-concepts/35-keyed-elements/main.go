package main

import "fmt"

func main() {
	names := [3]string{
		2: "Aj",
		0: "An",
		1: "Nik",
	}

	fmt.Printf("%q\n", names)

	//keyed elements

	names2 := [3]string{
		2: "Aj",
		0: "An",
		1: "Nik",
	}
	fmt.Printf("%q\n", names2)

	names3 := [...]string{
		5: "Aj",
		"Ansad",
		2: "Nik",
	}
	fmt.Printf("%q\n", names3)

	const (
		aju = iota
		anu
		nikh
		rah
	)

	ages := [...]float64{
		aju:  35,
		rah:  40,
		anu:  30,
		nikh: 33,
	}
	fmt.Printf("aju is %g old\n", ages[aju])
	fmt.Printf("nikh is %g old\n", ages[nikh])

	fmt.Printf("rah is %g old\n", ages[3])

	fmt.Printf("%#v\n", ages)

}
