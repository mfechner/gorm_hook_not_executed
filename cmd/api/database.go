package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm-hook-called-only-once/internal/types"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("testdb.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database successfully")

	log.Println("Running migrations...")
	err = db.AutoMigrate(&types.Domain{}, &types.Mailbox{})
	if err != nil {
		return nil, err
	}
	// Attach logger
	db.Logger = logger.Default.LogMode(logger.Info)
	return db, nil
}

func CloseDatabase(db *gorm.DB) error {
	sql, err := db.DB()
	if err != nil {
		return err
	}
	err = sql.Close()
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) SeedDatabase() {
	app.db.Exec("DELETE FROM domain")
	app.db.Exec("DELETE FROM mailbox")

	log.Println("Add Domains")
	domainTest1 := types.Domain{
		ID:     1,
		Domain: "test1",
	}
	app.db.Save(&domainTest1)
	domainTest2 := types.Domain{
		ID:     2,
		Domain: "test2",
	}
	app.db.Save(&domainTest2)

	log.Println("Add Mailboxes")
	mailboxTest1 := types.Mailbox{
		ID:       1,
		Name:     "Test1",
		DomainID: 1,
		Domain:   domainTest1,
	}
	app.db.Save(&mailboxTest1)
	mailboxTest2 := types.Mailbox{
		ID:       2,
		Name:     "Test2",
		DomainID: 1,
		Domain:   domainTest2,
	}
	app.db.Save(&mailboxTest2)

}
