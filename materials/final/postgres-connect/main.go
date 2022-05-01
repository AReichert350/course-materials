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
// \connect store
// create table transactions(ccnum varchar(32), date date, amount money, cvv char(4), exp date);
// insert into transactions(colName) values ('value');
// select * from transactions;