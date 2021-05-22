package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Contact struct {
	ID          int
	First_name  string
	Last_name   string
	Middle_name string
	Phone       string
	Email       string
	Address     string
	Created     time.Time
}
