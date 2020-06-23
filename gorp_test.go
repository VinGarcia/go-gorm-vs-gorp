package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

var gorpdb *gorp.DbMap

func init() {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres port=9999 dbname=test sslmode=disable")
	checkErr(err, "gorp could not open postgres db on localhost:9999")

	gorpdb = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	gorpdb.AddTableWithName(Model{}, "models").SetKeys(true, "id")
	checkErr(err, "gorp: create tables failed")
}

func gorpCleanup() {
	gorpdb.Exec("DELETE FROM models")
}

func BenchmarkGorp(b *testing.B) {
	var models []Model

	b.Run("Insert", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			m := NewModel()
			err := gorpdb.Insert(m)
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}

			if len(models) < 100 {
				models = append(models, *m)
			}
		}
	})

	b.Run("Update", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			m := models[n%len(models)]
			m.Name = fmt.Sprint("Updated Name ", n)
			m.Title = fmt.Sprint("Updated Title ", n)
			m.Counter++
			count, err := gorpdb.Update(&m)
			if err != nil || count == 0 {
				fmt.Println(count, err)
				b.FailNow()
			}
		}
	})

	b.Run("Read", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			m := models[n%len(models)]
			db := gormdb.Find(&m)
			if db.Error != nil {
				fmt.Println(db.Error)
				b.FailNow()
			}
		}
	})

	gorpCleanup()
}

func NewModel() *Model {
	return &Model{
		Name:    "Orm Benchmark",
		Title:   "Just a Benchmark for fun",
		Fax:     "99909990",
		Web:     "http://blog.milkpod29.me",
		Age:     100,
		Counter: 1000,
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg+":", err.Error())
		os.Exit(1)
	}
}
