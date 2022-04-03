package main

import (
	"log"
	"os"
	"hscan/hscan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <filename.txt>")
	}

	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	//TODO - Try to find these (you may or may not based on your password lists)
	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0"
	hscan.GuessSingle(drmike1, "Top304Thousand-probable-v2.txt")
		// drmike1 is p@ssword
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032"
	hscan.GuessSingle(drmike2, "Top304Thousand-probable-v2.txt")
		// drmike2 is letmein

	// NON CODE - TODO
	// Download and use bigger password file from: https://weakpass.com/wordlist/tiny  (want to push yourself try /small ; to easy? /big )

	//TODO Grab the file to use from the command line instead; look at previous lab (e.g., #3 ) for examples of grabbing info from command line
	// ./main wordlist.txt
	// var file = "wordlist.txt"
	var file = os.Args[1]

	hscan.GuessSingle(md5hash, file)
	hscan.GuessSingle(sha256hash, file)
	hscan.GenHashMaps(file)
	hscan.GetSHA(sha256hash)
	hscan.GetMD5(md5hash)
}
