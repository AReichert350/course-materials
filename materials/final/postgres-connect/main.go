// https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart

// apt-cache search postgres
// sudo apt-get install postgresql
// sudo service postgresql status

// uses port 5432

// sudo -u postgres psql
// create database store;
// \connect stores
// create table transactions(ccnum varchar(32), date date, amount money, cvv char(4), exp date);