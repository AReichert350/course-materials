// Codeanywhere:
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

// To connect:
// mysql --help | grep user
// mysql -u root -p
// quit

// Docker:
// docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
// docker run -it --link some-mysql:mysql --rm mysql sh -c 'exec mysql -h "$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'

// create DATABASE store;
// show DATABASES;
// use store;
// create table transactions(ccnum varchar(32), date date, amount float(7,2), cvv char(4), exp date);
// show tables;
// insert into transactions(colName) values ('value');
// select * from transactions;

// go get github.com/go-sql-driver/mysql

// For this file:
// go build main.go
// ./main

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/store")
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
