package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Contact struct {
	ID         int
	FirstName  string
	LastName   string
	MiddleName string
	Phone      string
	Email      string
	Address    string
	Created    time.Time
}
