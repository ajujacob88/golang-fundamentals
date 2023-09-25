//previous log parsing program rewritten using struct

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	//sum     map[string]int
	sum     map[string]result //total visits per domain
	domains []string          //unique domain names
	total   int               //total visits to all domains
}

func main() {

	p := parser{
		//sum: make(map[string]int),
		sum: make(map[string]result),
	}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v\n", fields)
			return
		}

		domain := fields[0]

		if _, ok := p.sum[domain]; !ok { //this is to add it to the slice, store only the unique names
			p.domains = append(p.domains, domain)
		}

		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			fmt.Printf("Wrong input: %v, err is %v\n", fields[1], err)
			return
		}

		p.total += visits
		//p.sum[domain].visits += visits
		// r := result{
		// 	domain: domain,
		// 	visits: visits + p.sum[domain].visits,
		// }
		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}

	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 50))

	//but usually we doesnt loop over the map, because map is not designed for that, so store it in a slice and then print
	// for domain, visits := range sum {
	// 	fmt.Printf("%-30s %10d\n", domain, visits)
	// }

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

}
