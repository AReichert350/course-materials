// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: Scans ports 1-1024 of scanme.nmap.org and reports which ones are open and which are closed
// Adeline Reichert
// 2/3/22
// COSC 4010-01 Lab 2
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int  // notice the capitalization here. access limited!
var closedports []int


func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		// conn, err := net.Dial("tcp", address) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		conn, err := net.DialTimeout("tcp", address, 1 * time.Second)
		if err != nil { 
			results <- 0
			closedports = append(closedports, p)
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
func PortScanner() (int, int) {  
	start := time.Now()	// To help tune for TODO 4
	ports := make(chan int, 1024)   // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d,open\n", port)	// Changed to comma separated value format
	}
	for _, port := range closedports {
		fmt.Printf("%d,closed\n", port)
	}
	finished := time.Now()	// To help tune for TODO 4
	fmt.Printf("Time elapsed: %v\n", finished.Sub(start))	// To help tune for TODO 4

	return len(openports), len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
