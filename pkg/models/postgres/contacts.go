package postgres

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx"
	"github.com/saepiae/contact/pkg/models"
)

type ContactModel struct {
	ConnPool *pgx.ConnPool
}

func (m *ContactModel) Insert(firstName string, lastName string, middleName string, phone string, email string, address string) (int, error) {
	var row int
	stmt := `INSERT INTO contact_table (first_name, last_name, middle_name, phone, email, address, created) VALUES 
	($1, $2, $3, $4, $5, $6, now()) RETURNING id`
	err := m.ConnPool.QueryRow(stmt, firstName, lastName, middleName, phone, email, address).Scan(&row)
	if err != nil {
		return 0, err
	}
	return int(row), nil
}

func (m *ContactModel) Get(id int) (*models.Contact, error) {
	result := &models.Contact{}
	stmt := `select id, first_name, last_name, middle_name, phone, email, address, created from contact_table where id = $1`
	row := m.ConnPool.QueryRow(stmt, id)
	err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.MiddleName, &result.Phone, &result.Email, &result.Address, &result.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return result, nil
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
