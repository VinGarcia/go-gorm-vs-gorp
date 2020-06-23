package main

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

var gormdb *gorm.DB

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=postgres port=9999 dbname=test sslmode=disable")
	checkErr(err, "gorm could not connect to db")
	gormdb = db
}

func gormCleanup() {
	gormdb.Exec("DELETE FROM models")
}

func BenchmarkGorm(b *testing.B) {
	var models []Model

	b.Run("Insert", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			m := NewModel()

			db := gormdb.Create(&m)
			if db.Error != nil {
				fmt.Println(db.Error)
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
			db := gormdb.Save(&m)
			if db.Error != nil {
				fmt.Println(db.Error)
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

	gormCleanup()
}
