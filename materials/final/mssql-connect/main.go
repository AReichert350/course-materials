// Codeanywhere (gave up on codeanywhere before fully testing these commands):
// https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-ubuntu?view=sql-server-ver15

// wget -qO- https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -
// sudo add-apt-repository "$(wget -qO- https://packages.microsoft.com/config/ubuntu/16.04/mssql-server-2019.list)"
// sudo apt-get install mssql-server

// Docker:
// https://hub.docker.com/_/microsoft-mssql-server
// docker run --name some-mssql -p 1433:1433 -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=4010goBHG!' -d microsoft/mssql-server-linux

// docker run --name some-mssql -p 1433:1433 -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=4010goBHG!' -d mcr.microsoft.com/mssql/server
// docker ps
// docker exec -it some-mssql /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P '4010goBHG!'

// create database store;
// go
// use store;
// create table transactions(ccnum varchar(32), date date, amount decimal(7,2), cvv char(4), exp date);
// go
// insert into transactions(ccnum, date, amount, cvv, exp) values ('4444333322221111', '2019-01-05', 100.12, '1234', '2020-09-01');
// insert into transactions(ccnum, date, amount, cvv, exp) values ('4444123456789012', '2019-01-07', 2400.18, '5544', '2021-02-01');
// insert into transactions(ccnum, date, amount, cvv, exp) values ('4465122334455667', '2019-01-29', 1450.87, '9876', '2020-06-01');
// go
// select * from transactions;
// go

// go get github.com/denisenkom/go-mssqldb

// For this file:
// go build main.go
// ./main

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// Establish a connection to the DB with sql.Open()
	// First param: which driver to use
	// Second param: connection string

	db, err := sql.Open("sqlserver", "sqlserver://sa:4010goBHG!@127.0.0.1:1433?database=store")
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	var (
		ccnum, date, cvv, amount, exp string
	)
	// Pass a MSSQL statement to db.Query()
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
		fmt.Println(ccnum, date, amount, cvv, exp)
	}
	if rows.Err() != nil {
		log.Panicln(err)
	}
}
