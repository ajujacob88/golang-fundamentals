//go:generate wire
package main

import "fmt"

type Greeter struct {
	message string
}

func NewGreeter(message string) *Greeter {
	return &Greeter{
		message: message,
	}
}

func (g *Greeter) Greet() {
	fmt.Println(g.message)
}

func main() {
	greeter := NewGreeter("Hello, world!")
	greeter.Greet()
}
