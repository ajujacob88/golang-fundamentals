package main

import "fmt"

//so why we might wanna use pointers..? This is because they can work acrross functions or even accross packages.. Ther is only a single memory, so every function can see it. if we share an address to a function it can retrieve and change the value that is stored in that address..this is also why pointers can be dangerous too..This is because any part of the code can change any value if they have got the address of that value.. However go is not that dangerous as the other languages with pointers because go automatically manages the memory..

func main() {
	var (
		P       *int
		counter int
		V       int
	)

	fmt.Println(P)
	fmt.Printf("P is nil and its address is %p\n", P)
	P = &counter
	fmt.Println(P)
	fmt.Printf("Now p stores the address of the counter variable which is int %p\n", P)

	//also we can compare
	if P == &counter {
		fmt.Printf("P is equal to counter's address %p == %p\n", P, &counter)
	}

	//now puting some value to counter variable

	counter = 150
	fmt.Printf("Counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("*P: %-13d   P: %-13p  addr of P: %-13p\n", *P, P, &P)

	V = *P
	fmt.Printf("V: %-13d addr of V: %-13p\n", V, &V)

	fmt.Println("\n------CHange counter--------")

	counter = 333
	fmt.Printf("Counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("*P: %-13d   P: %-13p  addr of P: %-13p\n", *P, P, &P)
	fmt.Printf("V: %-13d addr of V: %-13p\n", V, &V)

	fmt.Println("\n------CHange counter in passVal() function ---------")

	passVal(counter)
	fmt.Printf("Counter: %-13d addr: %-13p\n", counter, &counter)

	fmt.Println("\n------CHange counter in passValPtr() function - a func can change a value belong to another func (or scope) indirectly through a pointer--------")

	passValPtr(&counter)
	fmt.Printf("Now Counter: %-13d addr: %-13p\n", counter, &counter)

}

func passVal(n int) int {
	n += 105
	fmt.Println("n is ", n)
	return n

}

func passValPtr(pn *int) {
	*pn += 215
	fmt.Println("now pn is:", *pn)

}
