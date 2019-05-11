package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/twinj/uuid"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NameSpaceUUID -
var NameSpaceUUID *uuid.UUID

func init() {
	NameSpaceUUID, _ = uuid.Parse("7311c711-03bd-4ad7-b639-976d2e849edb")
}

// Manager -
type Manager struct {
	*gorm.DB
}

// NewPostgres initialize with Postgres
func NewPostgres(host string, port int, user, dbname, password string, useSSL bool) *Manager {
	// connectString = "host=myhost port=myport user=gorm dbname=gorm password=mypassword sslmode=enable"
	sslMode := "enable"
	if !useSSL {
		sslMode = "disable"
	}
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslMode,
	))
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	return &Manager{db}
}

// NewSqlite initialize with SQLite
func NewSqlite(fileName string) *Manager {
	db, err := gorm.Open("sqlite3", fileName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	return &Manager{db}
}

// HandleError wraps postgres errors
// FIXME: give hints on which errors are expected / unexpected?
// see https://github.com/lib/pq/blob/master/error.go for the list of error codes
func (m *Manager) HandleError(db *gorm.DB) error {
	if db.Error == nil {
		return nil
	}
	err := db.Error
	switch err.(type) {
	case *pq.Error:
		e := err.(*pq.Error)
		return fmt.Errorf("I'm sorry Dave, I'm afraid I can't do that: %s", e.Detail)
	}
	return err
}
