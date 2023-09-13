package main

import (
	"fmt"
	"os"
	"strings"
)

const str = "Hi, good morning Today is Monday and good morning"

func main() {
	words := strings.Fields(str)

	query := os.Args[1:]

labelAju:
	for _, q := range query {
		//search:
		for i, w := range words {
			// switch q {
			// case "Hi,", "is", "and":
			// 	break search   //we can use labelled break in switch also
			// }

			if q == w {
				fmt.Printf("#%-2d:  %q\n", i+1, q)
				break labelAju //this is the labelled break, here the break will exit from the parent loop where label is placed, not from the nested loop,, if there is no label, then it will exit from the current nested loop here
				//continue labelAju  //similarly labelled continue
			}

		}
	}

	//also goto is also in go, but rarley used

	goto labelAju
}
