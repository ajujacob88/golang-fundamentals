package main

//you can attach method to any types like integer, float or structs or arrays or maps or functions etc, not limited to structs...

func main() {
	// mobydick := book{
	// 	title: "Moby Dick",
	// 	price: 525,
	// }

	// gta := game{
	// 	title: "San Andreas",
	// 	price: 2000,
	// }

	// nfs := game{
	// 	title: "NFS",
	// }

	var (
		//mobydick = book{title: "Moby Dick", price: 525}
		gta = game{title: "San Andreas", price: 2000}
		nfs = game{title: "NFS"}
	)

	var items []*game

	items = append(items, &gta, &nfs)

	// for _, it := range items {
	// 	it.print()
	// }

	my := list(items)
	my.print()

	my = nil
	my.print()

	mobydick := book{title: "Moby Dick", price: 525}
	abc := book{title: "abc", price: 230}
	xyz := book{title: "xfdg", price: 100}

	var bookcollect []*book
	bookcollect = append(bookcollect, &mobydick, &abc, &xyz)

	mybooks := booklist(bookcollect)
	mybooks.print()

	mybooks.discount()
	mybooks.print()

}
