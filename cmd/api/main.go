package main

import (
	"log"

	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer func() {
		err := CloseDatabase(db)
		if err != nil {
			log.Fatal("Cannot close database")
		}
	}()

	app := &Application{db: db}

	app.SeedDatabase()

}
