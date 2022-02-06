package main

import "bhg-scanner/scanner"

func main(){
	var portRange []int
	for i := 1; i <= 65535; i++ {
		portRange = append(portRange, i)
	}
	// Currently configured to scan ports 1-1024
	scanner.PortScanner(portRange[0:1024], "scanme.nmap.org")
}