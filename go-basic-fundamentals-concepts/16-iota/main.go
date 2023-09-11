package main

import "fmt"

//iota is a built-on constant generatorwhich generates ever increasing umbers. iota starts from 0 and  will keeps increasing itself

func main() {
	const (
		sunday = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
		aju
	)
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday, aju)

	const (
		january = iota
		february
		march
	)
	fmt.Println(january, february, march)

	const (
		aj = iota + 25
		an
		nik
	)
	fmt.Println(aj, an, nik)

	const (
		a = -(10 + iota)
		_
		_
		b
		c
		_
		d
	)
	fmt.Println(a, b, c, d)
}
