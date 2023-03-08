//concurrency and go routines - just basics
// https://www.youtube.com/watch?v=ZDNtUgWRDdk
// https://www.youtube.com/watch?v=V-0ifUKCkBI&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=54

package main

import (
	"fmt"
	"time"
)

func main() {
	go number("a")
	fmt.Print("hi")
	number("b")
	go number("c")
	number("d")
	fmt.Print("aju")
	number("c")
	number("d")
}
func number(str string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(str)
	}
}
