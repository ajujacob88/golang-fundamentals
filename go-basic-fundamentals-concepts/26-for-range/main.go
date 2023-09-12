package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args {
		fmt.Printf("%q\n", v)
	}

	for i, v := range os.Args { //i index variable and v value variable
		if i == 0 {
			continue
		}
		fmt.Printf("again %q\n", v)
	}

	for _, v := range os.Args[1:] {
		fmt.Printf("once more %q\n", v)
	}

	for i, _ := range os.Args[1:] {
		fmt.Printf("index is  %q\n", i)
	}
	//if only index is needed, then we can remove the value variable entirely
	for i := range os.Args[1:] {
		fmt.Printf("actually no need to use blank identifier here for value ,, index is  %q\n", i)
	}
}
