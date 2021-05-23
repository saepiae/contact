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

func (m *ContactModel) Insert(firstName string, lastName string, middleName sql.NullString, phone string, email sql.NullString, address sql.NullString) (int, error) {
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

func (m *ContactModel) Update(id int, firstName string, lastName string, middleName sql.NullString, phone string, email sql.NullString, address sql.NullString) (int, error) {
	stmt := `update contact_table
	set first_name  = $1,
		last_name   = $2,
		middle_name = $3,
		phone       = $4,
		email       = $5,
		address     = $6
	where id = $7
	RETURNING id;`
	var row int
	err := m.ConnPool.QueryRow(stmt, firstName, lastName, middleName, phone, email, address, id).Scan(&row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		}
		return 0, err
	}
	return int(row), nil
}

func (m *ContactModel) Delete(id int) (int, error) {
	stmt := `DELETE FROM contact_table WHERE id = $1 RETURNING id;`
	var row int
	err := m.ConnPool.QueryRow(stmt, id).Scan(&row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		}
		return 0, err
	}
	return int(row), nil
}

func (m *ContactModel) FindAll() ([]*models.Contact, error) {
	stmt := `select id, first_name, last_name, middle_name, phone, email, address, created from contact_table`
	rows, err := m.ConnPool.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*models.Contact
	for rows.Next() {
		contact := &models.Contact{}
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.MiddleName, &contact.Phone, &contact.Email, &contact.Address, &contact.Created)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (m *ContactModel) FindDublicates() ([]int, error) {
	var ids [1]int
	return ids[:], nil
}
