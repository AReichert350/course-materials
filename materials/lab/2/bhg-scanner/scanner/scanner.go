// Adeline Reichert
// 2/3/22
// COSC 4010-01 Lab 2

// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: Navigate out of the scanner folder and into the main folder.
//         From there, run 'go build' then './main'
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports, results chan int, targAdd string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", targAdd, p)
		// conn, err := net.Dial("tcp", address) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		conn, err := net.DialTimeout("tcp", address, 1 * time.Second)
		if err != nil { 	// got some error e.g., closed/filtered
			results <- -1 * p
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(rangeToScan []int, targAdd string) (int, int, map[int]string) {
	//TODO 3 : ADD closed ports; currently code only tracks open ports
	// Moved into port scanner so code can be run multiple times with "fresh"
	// instances of these variables
	var openports []int  // notice the capitalization here. access limited!
	var closedports []int

	ports := make(chan int, len(rangeToScan)) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, targAdd)
	}

	go func() {
		for _, port := range rangeToScan {
			ports <- port
		}
	}()

	for i := 0; i < len(rangeToScan); i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
		} else if port < 0 {
			closedports = append(closedports, -1 * port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	portMap := map[int]string{}

	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d,open\n", port)	// Changed to comma separated value format
		portMap[port] = "open"
	}
	for _, port := range closedports {
		fmt.Printf("%d,closed\n", port)
		portMap[port] = "closed"
	}

	return len(openports), len(closedports), portMap // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
