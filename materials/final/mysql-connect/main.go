// apt-cache search mysql | more
// sudo apt-get install mysql-server
// mysql -> 3306, can change in config file: /etc/mysql/mysql.conf.d/
//   Changed to use port 3226
// sudo vim mysqld.cnf
// ls -ltr

// sudo service mysql status
// sudo service mysql restart
//  "                 stop
//  "                 start

// mysql --help | grep user
// mysql -u root -p
// quit

// create DATABASE store;
// show DATABASES;
// use store;
// create table transactions(ccnum varchar(32), date date, amount decimal(7,2), cvv char(4), exp date);
// show tables;

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:102218@tcp(127.0.0.1:3226)/store")
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	var (
		ccnum, date, cvv, exp string
		amount                float32
	)
	rows, err := db.Query("SELECT ccnum, date, amount, cvv, exp FROM transactions")
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()
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