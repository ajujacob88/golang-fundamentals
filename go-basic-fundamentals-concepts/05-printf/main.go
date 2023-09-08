// printf - prints formatted output

package main

import "fmt"

func main() {
	brand := "Google"
	qty := 50
	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand) //%s means string, used only with string values
	fmt.Printf("%d\n", 20)    //placeholder %d is for decimal / integer value
	fmt.Printf("%f\n", 20.65) //placeholder %f is for float value
	fmt.Printf("%t\n", false) //placeholder %t is for bool value
	fmt.Printf("%T\n", brand) // this placeholder is for type
	//above verbs are the type safe verbs, because go warn us if that verb is assigned to a different type(sothat we tend to know we are using it correctly), but below %v is type unsafe, because the type is automatically interpreted

	fmt.Printf("%v\n", qty)   //%v is for any value
	fmt.Printf("%v\n", brand) //%v is for any value

	fmt.Printf("Total: %d, outof: %d \n", 80, 100)

	//Escape sequences
	fmt.Println("aju\"jacob\" \\n is escape charactoe \n escape")

	fmt.Println()
	//printf index.. printf index starts with not 0, instead 1.. with the help of indexing we can pass these variables multiple times
	fmt.Printf("the brand is %v and qty is %v, sorry the qty is %[2]v and price is %[3]v and brand is %[1]v \n", brand, qty, 25000)

	fmt.Println()
	//for precisions
	decimal := 99.123456789
	fmt.Printf("%.0f\n", decimal)
	fmt.Printf("%.1f\n", decimal)
	fmt.Printf("%.2f\n", decimal)
	fmt.Printf("%.3f\n", decimal)

	fmt.Printf("%f\n", decimal)
	fmt.Printf("%v\n", decimal)
	fmt.Printf("%g\n", decimal) //%g will eliminate unnecessary fractional parts

}
