package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if(len(sourceHash) == 32) {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				return password
			}
		} else if (len(sourceHash) == 64) {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				return password
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ""
}

func GenHashMaps(filename string) {

	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	shalookup = make(map[string]string)
    md5lookup = make(map[string]string)

	// Without goroutines
	// fmt.Print("Generating hash maps without using goroutines")
	// for scanner.Scan() {
	// 	password := scanner.Text()
	// 	// Hash to MD5 and save
	// 	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	// 	md5lookup[hash] = password
	// 	// Hash to SHA-256 and save
	// 	hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	// 	shalookup[hash] = password
	// }

	// With goroutines
	// Used this website to help with goroutines:
	// https://yourbasic.org/golang/wait-for-goroutines-waitgroup/#:~:text=WaitGroup%20waits%20for%20a%20group%20of%20goroutines%20to%20finish.&text=First%20the%20main%20goroutine%20calls,and%20call%20Done%20when%20finished.
	fmt.Print("Generating hash maps using goroutines\n")
	var wg sync.WaitGroup
	for scanner.Scan() {
		password := scanner.Text()
		wg.Add(2)
		go func() {
			// Hash to MD5 and save
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			md5lookup[hash] = password
			wg.Done()
		}()
		go func() {
			// Hash to SHA-256 and save
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			shalookup[hash] = password
			wg.Done()
		}()
		wg.Wait()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
		// For Top304Thousand-probable-v2.txt, it takes 1.212 seconds to generate the maps without goroutines and
        // 5.012 seconds to generate the maps using goroutines.
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
		// For Top304Thousand-probable-v2.txt, it takes 3.99E-6 seconds per password when not using goroutines and
		// 1.65E-5 seconds per password when using goroutines.
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		fmt.Print("Password found in shalookup: ", password, "\n")
		return password, nil

	} else {
		fmt.Print("Password for hash ", hash, " not found in shalookup\n")
		return "", errors.New("password does not exist")

	}
}

//TODO
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		fmt.Print("Password found in md5lookup: ", password, "\n")
		return password, nil
	} else {
		fmt.Print("Password for hash ", hash, " not found in md5lookup\n")
		return "", errors.New("password does not exist")
	}
}
