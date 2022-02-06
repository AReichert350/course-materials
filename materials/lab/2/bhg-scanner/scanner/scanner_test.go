package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){
	var portRange []int
	for i := 1; i <= 65535; i++ {
		portRange = append(portRange, i)
	}
    numOpen, _, _ := PortScanner(portRange[0:1024], "scanme.nmap.org") // Function returns both number of open and number of closed ports
    want := 1 // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	          //consider what would happen if you parameterize the portscanner address and ports to scan

    if numOpen < want {
        t.Errorf("got %d, wanted at least %d", numOpen, want)
    }
}

func TestTotalPortsScanned(t *testing.T){
	var portRange []int
	for i := 1; i <= 65535; i++ {
		portRange = append(portRange, i)
	}
	portRange = portRange[0:1024]
    numOpen, numClosed, _ := PortScanner(portRange, "scanme.nmap.org") // Functions returns both number of open and number of closed ports
	got := numOpen + numClosed
    want := len(portRange) // default value; consider what would happen if you parameterize the portscanner ports to scan

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}

func TestPort22Open(t *testing.T){
	var portRange []int
	for i := 1; i <= 65535; i++ {
		portRange = append(portRange, i)
	}
    _, _, portMap := PortScanner(portRange[21:27], "scanme.nmap.org") // Function returns both number of open and number of closed ports

    if portMap[22] != "open" {
        t.Errorf("got %s, wanted open", portMap[22])
    }
}


