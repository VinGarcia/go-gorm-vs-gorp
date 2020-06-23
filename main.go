package main

type Model struct {
	ID      int    `orm:"auto" gorm:"primary_key" db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Counter int64  `db:"counter"`
}

var createTableQuery = `
CREATE TABLE models(
  id SERIAL PRIMARY KEY,
	name varchar(255),
	title varchar(255),
	fax varchar(255),
	web varchar(255),
	age int,
	counter int
)
`
