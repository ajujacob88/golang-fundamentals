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

	//converting the above functional style of programming to methods style,,

	mobydick := book{
		title: "Moby Dick",
		price: 525,
	}
	// mobydick.printBook()
	mobydick.print()
	book.print(mobydick) //this is same as mobydick.print()..this is called method expressions..This allows you to call a method through a type.. this works becasue behind the scenes a method exactly looks and works like a function., the print method looks like this:- func print(b book){}..behind the scenes, a method is a function with recceiver as its first parameter., so we can call the dunction in this way also.. so here  mobydick.print() auomatically passes mobydick as an argument to print(), while in book.print(mobydick) manually passes mobydick as an argument to print() by ourself,, therefore in bothcases the print() function gets a copy of the mobydick,, either way is same, but we usually use this way like mobydick.print()

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
