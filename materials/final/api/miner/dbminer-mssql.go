package miner

import (
	"database/sql"
	"fmt"
	"log"

	// "os"

	// "mssqlDBMiner/dbminer"

	_ "github.com/denisenkom/go-mssqldb"
)

type MSSQLMiner struct {
	Host string
	Db   sql.DB
}

func MSSQLNew(host string) (*MSSQLMiner, error) {
	m := MSSQLMiner{Host: host}
	err := m.connect()
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *MSSQLMiner) connect() error {
	// Connect to the overall mssql instance
	db, err := sql.Open("sqlserver", fmt.Sprintf("sqlserver://sa:4010goBHG!@%s:1433", m.Host))
	if err != nil {
		log.Panicln(err)
	}
	m.Db = *db
	return nil
}

func (m *MSSQLMiner) GetSchema() (*Schema, error) {
	var s = new(Schema)

	// Get the names of all the DBs at the IP address given
	// Only query for user-defined DBs (exclude the default mssql DBs)
	sqlQuery := `SELECT name 
	             FROM sys.databases 
							 WHERE name not in ('master', 'tempdb', 'model', 'msdb')`
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
		// log.Println("Building schema for DB " + name)
		// Connect to the DB
		dbConn, err := sql.Open("sqlserver", fmt.Sprintf("sqlserver://sa:4010goBHG!@%s:1433?database=%s", m.Host, name))
		if err != nil {
			log.Panicln(err)
		}
		// Get the DB's tables
		// https://chartio.com/resources/tutorials/sql-server-list-tables-how-to-show-all-tables/
		sqlQuery = `SELECT name
		            FROM sysobjects
								WHERE xtype = 'U'`
		tables, err := (*dbConn).Query(sqlQuery)
		if err != nil {
			return nil, err
		}
		defer tables.Close()

		var dbInitialized = false
		// var prevtable string
		var db Database
		var table Table
		for tables.Next() {
			if !dbInitialized {
				db = Database{Name: name, Tables: []Table{}}
				dbInitialized = true
			}

			var currtable string
			if err := tables.Scan(&currtable); err != nil {
				return nil, err
			}
			table = Table{Name: currtable, Columns: []string{}}
			// log.Println("collecting col names for table " + currtable)

			// Get the table's columns
			// https://www.mytecbits.com/microsoft/sql-server/list-of-column-names
			sqlQuery = fmt.Sprintf(`SELECT column_name
															FROM information_schema.columns
															WHERE table_name = '%s'`, currtable)
			cols, err := (*dbConn).Query(sqlQuery)
			if err != nil {
				return nil, err
			}
			defer cols.Close()

			for cols.Next() {
				var currcol string
				if err := cols.Scan(&currcol); err != nil {
					return nil, err
				}
				// log.Println("found col name " + currcol)
				table.Columns = append(table.Columns, currcol)
			}
			db.Tables = append(db.Tables, table)
		}
		if dbInitialized {
			s.Databases = append(s.Databases, db)
			dbInitialized = false
		}
		if err := tables.Err(); err != nil {
			return nil, err
		}
	}

	// fmt.Println("Returning schema:")
	// fmt.Println(s)
	return s, nil
}

func MSSQLMain(ip_addr string) []string {
	mm, err := MSSQLNew(ip_addr)
	if err != nil {
		panic(err)
	}
	defer mm.Db.Close()

	mineResults, err := Search(mm)
	if err != nil {
		panic(err)
	}

	return mineResults
}
