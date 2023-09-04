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

	speed, distance, currentTime := 150, 25, time.Now()
	fmt.Println(speed, distance, currentTime)

	speed, distance = distance, speed
	fmt.Println(speed, distance)
}
