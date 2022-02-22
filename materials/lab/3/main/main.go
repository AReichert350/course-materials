// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <S | Q>

package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"shodan/shodan"
)

// Notes from lab video:
// To accept additional command line args: 
//    change 2 to accept different num
//    for variable num of args, can use ||
//    <requiredArgName> [optionalArgName] in Usage line
// Then use os.Args[1], os.Args[2] wherever you need to use the entered value
// May have to typecast

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <S | Q>")
	}
	if os.Args[1] == "S" {
		HostSearch()
	} else if os.Args[1] == "Q" {
		Query()
	} else {
		log.Fatalln("Please enter 'S' to perform a Host Search or 'Q' to perform a query")
	}
}

func HostSearch() {
	var searchTerm string
	fmt.Println("Please enter your search term and press Enter.")
	fmt.Scanln(&searchTerm)

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	nextPage := "Y"
	page := 0

	for nextPage == "Y" {
		page++

		fmt.Printf(
			"Query Credits: %d\nScan Credits:  %d\n\n",
			info.QueryCredits,
			info.ScanCredits)

		hostSearch, err := s.HostSearch(searchTerm, page)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("Host Data Dump\n")
		for _, host := range hostSearch.Matches {
			fmt.Println("==== start ",host.IPString,"====")
			h,_ := json.Marshal(host)
			fmt.Println(string(h))
			fmt.Println("==== end ",host.IPString,"====")
		}

		fmt.Printf("IP, Port, City\n")

		for _, host := range hostSearch.Matches {
			fmt.Printf("%s, %d, %s\n", host.IPString, host.Port, host.Location.City)
		}

		fmt.Println("Press Y and Enter to get next page.")
		fmt.Scanln(&nextPage)
	}
}

func Query() {
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	nextPage := "Y"
	page := 0

	for nextPage == "Y" {
		page++

		fmt.Printf(
			"Query Credits: %d\nScan Credits:  %d\n\n",
			info.QueryCredits,
			info.ScanCredits)

		query, err := s.Query(page)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("Query Data Dump\n")
		for _, q := range query.Matches {
			fmt.Println("==== start ",q.Title,"====")
			h,_ := json.Marshal(q)
			fmt.Println(string(h))
			fmt.Println("==== end ",q.Title,"====")
		}

		//TODO: Left off here
		fmt.Printf("Title, Votes\n")

		for _, q := range query.Matches {
			fmt.Printf("%s, %d\n", q.Title, q.Votes)
		}

		fmt.Println("Press Y and Enter to get next page.")
		fmt.Scanln(&nextPage)
	}
}