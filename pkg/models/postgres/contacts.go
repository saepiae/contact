package postgres

import (
	"github.com/jackc/pgx"
	"github.com/saepiae/contact/pkg/models"
)

type ContactModel struct {
	ConnPool *pgx.ConnPool
}

func (m *ContactModel) Insert(first_name string, last_name string, middle_name string, phone string, email string, address string) (int, error) {
	return 0, nil
}

func (m *ContactModel) Get(id int) (*models.Contact, error) {
	return nil, nil
}

func (m *ContactModel) Update(id int, first_name string, last_name string, middle_name string, phone string, email string, address string) (int, error) {
	return 0, nil
}

func (m *ContactModel) Delete(id int) (int, error) {
	return 0, nil
}

func (m *ContactModel) FindAll() ([]*models.Contact, error) {
	return nil, nil
}

func (m *ContactModel) FindDublicates() ([]int, error) {
	var ids [1]int
	return ids[:], nil
}
