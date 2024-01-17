package config

import (
	"github.com/humbertodlacerda/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/openings.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating ...")
		// Create the database file and directory
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(
		&schemas.Opening{},
		&schemas.Movement{},
	)
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
	}

	return db, nil
}
