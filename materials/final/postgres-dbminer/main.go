// To run:
// go build main.go
// ./main 127.0.0.1

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"postgresDBMiner/dbminer"

	_ "github.com/lib/pq"
)

type PostgresMiner struct {
	Host string
	Db   sql.DB
}

func New(host string) (*PostgresMiner, error) {
	m := PostgresMiner{Host: host}
	err := m.connect()
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *PostgresMiner) connect() error {
	// Connect to the overall postgres instance
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=5432 user=postgres password=goBHG sslmode=disable", m.Host))
	if err != nil {
		log.Panicln(err)
	}
	m.Db = *db
	return nil
}

func (m *PostgresMiner) GetSchema() (*dbminer.Schema, error) {
	var s = new(dbminer.Schema)

	// Get the names of all the DBs at the IP address given
	sqlQuery := `SELECT datname 
	             FROM pg_database
							 WHERE datname not in ('template0')`
	dbnames, err := m.Db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer dbnames.Close()

	var (
		name string
	)

	// For each database, get its schema
	for dbnames.Next() {
		err := dbnames.Scan(&name)
		if err != nil {
			log.Panicln(err)
		}
		// fmt.Println("Building schema for DB " + name)
		// Connect to the DB
		dbConn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=5432 user=postgres password=goBHG dbname=%s sslmode=disable", m.Host, name))
		if err != nil {
			log.Panicln(err)
		}
		// Get the DB's schema
		sqlQuery = `SELECT table_name, column_name
								FROM information_schema.columns
								WHERE table_schema not in ('information_schema', 'pg_catalog')
								ORDER BY table_name`
		schema_rows, err := (*dbConn).Query(sqlQuery)
		if err != nil {
			return nil, err
		}
		defer schema_rows.Close()

		var dbInitialized = false
		var prevtable string
		var db dbminer.Database
		var table dbminer.Table
		for schema_rows.Next() {
			if !dbInitialized {
				db = dbminer.Database{Name: name, Tables: []dbminer.Table{}}
				dbInitialized = true
			}
			var currtable, currcol string
			if err := schema_rows.Scan(&currtable, &currcol); err != nil {
				return nil, err
			}

			if currtable != prevtable {
				if prevtable != "" {
					db.Tables = append(db.Tables, table)
				}
				table = dbminer.Table{Name: currtable, Columns: []string{}}
				prevtable = currtable
			}
			table.Columns = append(table.Columns, currcol)
		}
		if dbInitialized {
			db.Tables = append(db.Tables, table)
			s.Databases = append(s.Databases, db)
			dbInitialized = false
		}
		if err := schema_rows.Err(); err != nil {
			return nil, err
		}
	}

	// fmt.Println("Returning schema:")
	// fmt.Println(s)
	return s, nil
}

// Identical to the mongo db miner
func main() {
	mm, err := New(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer mm.Db.Close()

	if err := dbminer.Search(mm); err != nil {
		panic(err)
	}
	return
}
