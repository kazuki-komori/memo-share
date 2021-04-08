package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/oklog/ulid"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewDB() (*SqlHandler, error) {

	db, err := gorm.Open("postgres", "postgres://admin:pass@db:5432/db?sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("failed to connect db= %w", err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler, nil
}

func CreateULID() (ULID string) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
