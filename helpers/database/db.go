package database

import (
	"Food/helpers/setting"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(DatabaseSetting setting.Database) (*gorm.DB, func()) {
	// dialector := postgres.New(postgres.Config{
	// 	DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
	// 		DatabaseSetting.Host,
	// 		DatabaseSetting.User,
	// 		DatabaseSetting.Password,
	// 		DatabaseSetting.Name),
	// })
	dialector := sqlite.Open("test.db")

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