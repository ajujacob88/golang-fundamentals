package main

import "fmt"

func main() {
	map1 := map[string]int{"aju": 32, "eldho": 40, "deepa": 35}

	mapList := []map[string]interface{}{ //if we use interface{} keyword, then we can insert any types like int or float or string etc
		{"name": "jobin", "age": 29},
		{"name": "kamal", "age": 21, "place": "mvpa", "Pin": 686673},
		{"name": "bibin", "age": 28},
	}
	fmt.Println(mapList[2])

	mapList1 := []map[interface{}]interface{}{
		{"name": "jobin", "age": 29},
		{"name": "kamal", "age": 21, "place": "mvpa", 2022: "ekm"},
		{"name": "bibin", "age": 28},
	}

	fmt.Println(mapList1[1])

	mapList2 := []map[string]int{
		{"name": 52, "age": 29},
		{"name": 62, "age": 21},
		{"name": 72, "age": 28},
	}

	map2 := []map[string]int{{"aju": 32, "eldho": 40}, {"deepa": 35, "eldho": 40}}
	fmt.Println(map1)
	fmt.Println(mapList)
	fmt.Println(mapList1)
	fmt.Println(mapList2)
	fmt.Println(map2)
}

/* from chatgpt - read aju
The Go programming language specification does not guarantee any specific order of iteration when ranging over a map.
Therefore, the order of the printed output may vary between different Go compilers or even between different runs of the same code.

However, in some Go compilers, including the official Go compiler and some third-party compilers, the implementation of the map type iterates over the keys in a sorted order by default.
This means that when you print the map, the keys will appear in alphabetical order.

It is important to note that the order of the keys is not guaranteed by the language specification and should not be relied upon in your code.
If you need to iterate over the map in a specific order, you should sort the keys explicitly before iterating over them.


To sort the keys of a map, you can follow these steps:

Extract the keys from the map into a slice.
Sort the slice using the sort package.
Iterate over the sorted slice, accessing the values in the original map using the keys.
Here's an example code snippet that demonstrates how to sort the keys of a map in Go:


import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

    // Step 1: Extract the keys from the map into a slice.
    var keys []string
    for k := range m {
        keys = append(keys, k)
    }

    // Step 2: Sort the slice using the sort package.
    sort.Strings(keys)

    // Step 3: Iterate over the sorted slice, accessing the values in the original map using the keys.
    for _, k := range keys {
        fmt.Printf("key=%s value=%d\n", k, m[k])
    }
}

In this example, the keys slice is created by iterating over the keys of the m map and appending them to the slice.
The sort.Strings function is then used to sort the keys slice in alphabetical order.
Finally, the sorted keys are iterated over using a for loop and the corresponding values are accessed from the original m map.

*/
