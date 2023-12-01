package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)
	speed, now = 100, time.Now()
	fmt.Println(speed, now)

	speed, distance, currentTime := 150, 250, time.Now()
	fmt.Println(speed, distance, currentTime)

	speed, distance = distance, speed
	n, err := fmt.Println(speed, distance)

	fmt.Println("n is", n, "err iss", err)
}
