package main

import "fmt"

func main() {
	//spendings 1st day = 100,800,450,200
	//spendings 2nd day = 5000,100
	//spendings 3rd day = 80,350,9000,250,550

	spendings := [][]int{
		{100, 800, 450, 200},
		{5000, 100},
		{80, 350, 9000, 250, 550},
	}
	fmt.Println(spendings)

	for i, daily := range spendings {
		fmt.Println("i:", i, "daily:", daily)
		var total int
		for _, v := range daily {
			total += v
		}
		fmt.Println("daily total:", total)
	}

	//another eg
	fmt.Println("another one")
	earnings := make([][]int, 0, 5)                   //earnings for each day, 5 days
	earnings = append(earnings, []int{100, 200, 300}) //here no need of ellipses... since its a multi dim slices, and we expect a slice to pass into this append
	earnings = append(earnings, []int{1100})
	earnings = append(earnings, []int{1500, 3000, 4000, 100})
	earnings = append(earnings, []int{2000, 1000}, []int{3500, 200}) //we can append all these values in one step as well,,, also we can create new append, but it will create new backarray as iam declared make capacity as 5, so in the background, more costly

	fmt.Println(earnings)
	for i, daily := range earnings {
		var total int
		for _, v := range daily {
			total += v
		}
		fmt.Printf("day %v, total earnings: %v\n", i+1, total)
	}

}
