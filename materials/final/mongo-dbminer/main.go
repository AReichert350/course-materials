// To run:
// go build main.go
// ./main 127.0.0.1

package main

import (
	"os"
	// "fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"mongoDBMiner/dbminer"
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

func (m *MongoMiner) GetSchema() (*dbminer.Schema, error) {
	var s = new(dbminer.Schema)

	dbnames, err := m.session.DatabaseNames()
	if err != nil {
		return nil, err
	}

	for _, dbname := range dbnames {
		db := dbminer.Database{Name: dbname, Tables: []dbminer.Table{}}
		collections, err := m.session.DB(dbname).CollectionNames()
		if err != nil {
			return nil, err
		}

		for _, collection := range collections {
			table := dbminer.Table{Name: collection, Columns: []string{}}

			// Allows you to unmarshal the structured data without needing
			// to know the structure of the data beforehand
			var docRaw bson.Raw
			// fmt.Println(dbname)
			// fmt.Println(collection)
			// fmt.Println()
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
		s.Databases = append(s.Databases, db)
	}
	return s, nil
}

// Expects the IP address of a MongoDB instance as an argument
func main() {

	mm, err := New(os.Args[1])
	if err != nil {
		panic(err)
	}
	// Search() calls getSchema()
	if err := dbminer.Search(mm); err != nil {
		panic(err)
	}
}
