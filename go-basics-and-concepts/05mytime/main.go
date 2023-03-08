package main

import (
	"fmt"
	"time"
)

func main() {
	//ref: https://youtu.be/xsAz24z4Hdg

	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))                 //this value is used for standard date formating
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //this value need to be used for format, this value can be found out from documentation. this is the syntax, read that documentation time part for similar time formats.

	//if we want to create the time based on the values we are entering
	createdDate := time.Date(1988, time.June, 4, 13, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	presentTime1 := time.Now().Nanosecond()
	fmt.Println(presentTime1)

	//we can build for windows / apple/linux using the command - go build for the current os and ' GOOS="linux" go build '  ,,for mac os - ' GOOS="darwin" go build '  BUT THIS IS NOT WORKING NOW
} //ref: https://youtu.be/9vxsO5vMTAw
