package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func NewDB() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "host=db port=5432 user=admin dbname=db password=pass")

	if err != nil {
		return nil, fmt.Errorf("failed to conntect db= %w", err)
	}
	return db, nil
}
