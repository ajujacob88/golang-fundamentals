package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
		total   int
	)
	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v\n", fields)
			return
		}

		domain := fields[0]

		if _, ok := sum[domain]; !ok { //this is to add it to the slice, store only the unique names
			domains = append(domains, domain)
		}

		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			fmt.Printf("Wrong input: %v, err is %v\n", fields[1], err)
			return
		}

		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 50))

	//but usually we doesnt loop over the map, because map is not designed for that, so store it in a slice and then print
	// for domain, visits := range sum {
	// 	fmt.Printf("%-30s %10d\n", domain, visits)
	// }

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

}
