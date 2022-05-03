// Codeanywhere:
// https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart

// apt-cache search postgres
// sudo apt-get install postgresql
// sudo service postgresql status

// uses port 5432

// To connect:
// sudo -u postgres psql

// Docker:
// docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=goBHG -d postgres
// docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres

// create database store;
// \list
// \connect store
// create table transactions(ccnum varchar(32), date date, amount money, cvv char(4), exp date);
// insert into transactions(colName) values ('value');
// select * from transactions;

// go get github.com/lib/pq

// For this file:
// go build main.go
// ./main

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	// Establish a connection to the DB with sql.Open()
	// First param: which driver to use
	// Second param: connection string

	// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=goBHG dbname=store sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	var (
		ccnum, date, cvv, amount, exp string
	)
	// Pass a PostgreSQL statement to db.Query()
	rows, err := db.Query("SELECT ccnum, date, amount, cvv, exp FROM transactions")
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
	// Loop through all the rows returned by db.Query()
	for rows.Next() {
		err := rows.Scan(&ccnum, &date, &amount, &cvv, &exp)
		if err != nil {
			log.Panicln(err)
		}
		// Slight data "cleaning"
		amount = amount[1:]
		amount = strings.ReplaceAll(amount, ",", "")
		fmt.Println(ccnum, date, amount, cvv, exp)
	}
	if rows.Err() != nil {
		log.Panicln(err)
	}
}
