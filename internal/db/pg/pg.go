package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(postgresDsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
