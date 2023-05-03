// return unique element from the string - interview question

package main

import "fmt"

func main() {

	var string1 = "aabbcddee"
	fmt.Println(uniqueElement(string1))
}

func uniqueElement(data string) string {
	flag := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] {
				flag++
			}
		}
		if flag == 1 {
			return string(data[i])
		}
		flag = 0
	}
	return "null"

}
