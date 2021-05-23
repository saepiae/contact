package models

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Contact struct {
	ID         int            `json:"id"`
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	MiddleName sql.NullString `json:"middleName"`
	Phone      string         `json:"phone"`
	Email      sql.NullString `json:"email"`
	Address    sql.NullString `json:"address"`
	Created    time.Time      `json:"created"`
}
