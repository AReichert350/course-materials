// sudo apt-get install mongodb
// go mod init ch7.go
// go mod tidy
// go build main.go

// sudo service mongodb status

// Our MongoDB is mapped to port 27017?

// To connect:
// mongo
// > db.transactions.insert([ {"name1" : "value1"} ])
// > db.transactions.find( {} )

// go get gopkg.in/mgo.v2

// Ch7 Notes:
// - MongoDB uses a JSON syntax for retrieving and manipulating data.
// - Unlike traditional SQL databases, MongoDB is schema-less, meaning it doesn't
//   follow a predefined, rigid rule system for organizing table data.

package main

import (
	"fmt"
	// "log"

	mgo "gopkg.in/mgo.v2"
)

// bson is binary JSON
type Transaction struct {
	CCNum      string  `bson:"ccnum"`
	Date       string  `bson:"date"`
	Amount     float32 `bson:"amount"`
	Cvv        string  `bson:"cvv"`
	Expiration string  `bson:"exp"`
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(session)
	}
	defer session.Close()

	results := make([]Transaction, 0)
	if err := session.DB("store").C("transactions").Find(nil).All(&results); err != nil {
		fmt.Println(err)
	}
	for _, txn := range results {
		fmt.Println(txn.CCNum, txn.Date, txn.Amount, txn.Cvv, txn.Expiration)
	}
}