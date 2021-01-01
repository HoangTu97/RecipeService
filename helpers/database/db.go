package database

import (
	"Food/helpers/setting"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(config setting.Database) (*gorm.DB, func()) {
	var dialector gorm.Dialector

	switch config.Type {
	case "sqlite3":
		dialector = sqlite.Open("test.db")
	case "postgres":
		dialector = postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
				config.Host,
				config.User,
				config.Password,
				config.Name),
		})
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	teardown := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
	return db, teardown
}
