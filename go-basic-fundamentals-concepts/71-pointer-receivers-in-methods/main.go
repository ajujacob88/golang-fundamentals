//go run .

package main

//summary
// a method with a value receiver always copy the receiver when it executes, so when we receive a large value, then again the data will be copied which takes too much memory as wellas execution time. But instead if we use pointer receivers, then the data is not copied, and working wth original data, so it will takes less memory and time of execution
//so in summary, when to use? use a pointer receiver when you want to change a receiver variable or use a pointer receiver when the received value is very large..

func main() {

	mobydick := book{
		title: "Moby Dick",
		price: 525,
	}

	gta := game{
		title: "San Andreas",
		price: 2000,
	}

	nfs := game{
		title: "NFS",
	}

	mobydick.print() //no need to pass any pointer here like (&mobydick).print() as the receiver is a pointer , because go automatically handles that,... ie when a method has a pointer receiver Go can take its address automatically
	//(&mobydick).print()
	mobydick.discount(20) //here since the discount method receiver is a pointer, the original value mobydick.discount is automaticallly changed from inside the discount menthod
	mobydick.print()

	gta.print()

	nfs.print()

}
