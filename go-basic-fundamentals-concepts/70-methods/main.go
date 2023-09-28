//go run .

package main

import "fmt"

//methods enhance type with additional behaviour

func main() {
	// mobydick := book{
	// 	title: "Moby Dick",
	// 	price: 525,
	// }
	// printBook(mobydick)

	// gta := game{
	// 	title: "San Andreas",
	// 	price: 2000,
	// }
	// printGame(gta)

	//converting the above functional style of programming to methods style

	mobydick := book{
		title: "Moby Dick",
		price: 525,
	}
	// mobydick.printBook()
	mobydick.print()

	gta := game{
		title: "San Andreas",
		price: 2000,
	}
	//gta.printGame()
	gta.print()

	discountedPrice := mobydick.discount(20)
	fmt.Println(discountedPrice)

	nfs := game{
		title: "NFS",
	}
	nfs.print()

}
