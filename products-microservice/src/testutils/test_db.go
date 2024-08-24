package testutils

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(migrations ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(migrations...)
	if err != nil {
		panic("failed to migrate models")
	}

	return db
}
