package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	connStr := "postgres://postgres:admin@localhost/postgres?sslmode=disable"

	db, err := gorm.Open("postgres", connStr)
	if err != nil{
		log.Fatal(err)
	}
	return db
}



