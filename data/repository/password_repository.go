package repository

import (
	"database/sql"
	"log"
	"password-bank-golang/service/model"
)

type IPasswordRepository interface {
}

type PasswordRepository struct {
	conn *sql.DB
}

func (repo *PasswordRepository) GetPasswordById(passId string) (model.Password, error) {
	query := "SELECT 1 FROM password WHERE id = $1"
	row := repo.conn.QueryRow(query, passId)

	pass, err := extractPasswordFromRow(row)
	if err != nil {
		return model.Password{}, err
	}

	return pass, nil
}

func (repo *PasswordRepository) GetAllPasswords() ([]model.Password, error) {
	query := "SELECT * FROM password"
	rows, err := repo.conn.Query(query)

	if err != nil {
		return nil, err
	}

	passwords, err := extractPasswordsFromRows(rows)
	if err != nil {
		return []model.Password{}, err
	}

	return passwords, nil
}

func (repo *PasswordRepository) CreatePassword(passModel model.Password, email string) error {
	query := "INSERT INTO password(password, type, account, service_name, notes, date, user_email) VALUES($1, $2, $3, $4, $5, $6, $7)"
	_, err := repo.conn.Exec(query, passModel.Password, passModel.Type, passModel.Account, passModel.ServiceName, passModel.Notes, passModel.Date, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *PasswordRepository) UpdatePassword(passModel model.Password) error {
	query := "UPDATE password SET password = $1, type = $2, account = $3, service_name = $4, notes = $5, date = $6 WHERE id = $7"
	_, err := repo.conn.Exec(
		query,
		passModel.Password,
		passModel.Type,
		passModel.Account,
		passModel.ServiceName,
		passModel.Notes,
		passModel.Date,
		passModel.Id,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *PasswordRepository) DeletePassword(id string) error {
	query := "DELETE FROM password WHERE id = $1"
	_, err := repo.conn.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func extractPasswordsFromRows(rows *sql.Rows) ([]model.Password, error) {
	var passwords []model.Password
	var id int32
	var password string
	var pType string
	var account string
	var serviceName string
	var notes string
	var date string

	for rows.Next() {
		if err := rows.Scan(&id, &password, &pType, &account, &serviceName, &notes, &date); err != nil {
			log.Println(err)
			return []model.Password{}, err
		} else {
			passwords = append(passwords, model.Password{
				Id:          id,
				Password:    password,
				Type:        pType,
				Account:     account,
				ServiceName: serviceName,
				Notes:       notes,
				Date:        date,
			})
			return passwords, nil
		}
	}
	return passwords, nil
}

func extractPasswordFromRow(row *sql.Row) (model.Password, error) {
	var passModel model.Password
	var id int32
	var password string
	var pType string
	var account string
	var serviceName string
	var notes string
	var date string

	if err := row.Scan(&id, &password, &pType, &account, &serviceName, &notes, &date); err != nil {
		log.Println(err)
		return model.Password{}, err
	} else {
		passModel = model.Password{
			Id:          id,
			Password:    password,
			Type:        pType,
			Account:     account,
			ServiceName: serviceName,
			Notes:       notes,
			Date:        date,
		}
		return passModel, nil
	}
}

func NewPasswordRepository(dbConn *sql.DB) IPasswordRepository {
	return &PasswordRepository{conn: dbConn}
}
