// To run:
// go build main.go
// ./main 127.0.0.1

// https://golangbyexample.com/functoin-different-package-go/#:~:text=The%20function%20in%20another%20package,first%20which%20contains%20that%20function.

package miner

import (

	// "fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "miner/dbminer"
)

type MongoMiner struct {
	Host    string
	session *mgo.Session
}

func New(host string) (*MongoMiner, error) {
	m := MongoMiner{Host: host}
	err := m.connect()
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *MongoMiner) connect() error {
	s, err := mgo.Dial(m.Host)
	if err != nil {
		return err
	}
	m.session = s
	return nil
}

func (m *MongoMiner) GetSchema() (*Schema, error) {
	var s = new(Schema)

	dbnames, err := m.session.DatabaseNames()
	if err != nil {
		return nil, err
	}

	for _, dbname := range dbnames {
		db := Database{Name: dbname, Tables: []Table{}}
		collections, err := m.session.DB(dbname).CollectionNames()
		if err != nil {
			return nil, err
		}

		for _, collection := range collections {
			table := Table{Name: collection, Columns: []string{}}

			// Allows you to unmarshal the structured data without needing
			// to know the structure of the data beforehand
			var docRaw bson.Raw
			// log.Printf(dbname)
			// log.Printf(collection)
			// fmt.Println()
			// For whatever reason, when running this via the API webpage, it
			// does not like the system.sessions table and crashes on the
			// .Find().One() line below.
			if collection != "system.sessions" {
				err := m.session.DB(dbname).C(collection).Find(nil).One(&docRaw)
				if err != nil {
					return nil, err
				}

				var doc bson.RawD
				if err := docRaw.Unmarshal(&doc); err != nil {
					if err != nil {
						return nil, err
					}
				}

				for _, f := range doc {
					table.Columns = append(table.Columns, f.Name)
				}
				db.Tables = append(db.Tables, table)
			}
		}
		s.Databases = append(s.Databases, db)
	}
	return s, nil
}

// Expects the IP address of a MongoDB instance as an argument
func MongoMain(ip_addr string) []string {
	mm, err := New(ip_addr)
	if err != nil {
		panic(err)
	}
	// Search() calls getSchema()
	mineResults, err := Search(mm)
	if err != nil {
		panic(err)
	}

	return mineResults
}
