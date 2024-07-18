package repository

import (
	"database/sql"
	"errors"
	"log"
)

type IUserRepository interface {
	CreateUser(userEmail string, userPassword string) error
	LoginUser(userEmail string, userPassword string) error
}

type UserRepository struct {
	conn *sql.DB
}

func (repo *UserRepository) CreateUser(userEmail string, userPassword string) error {
	exists, err := isUserExists(repo.conn, userEmail)

	if err != nil {
		log.Println(err)
		return err
	}

	if exists {
		return errors.New("user already exists")
	} else {
		query := "INSERT INTO public.user (email, password) VALUES($1, $2)"
		_, err = repo.conn.Exec(query, userEmail, userPassword)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}

func (repo *UserRepository) LoginUser(userEmail string, userPassword string) error {
	var userExists bool
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE email = $1 AND password = $2)"

	err := repo.conn.QueryRow(query, userEmail, userPassword).Scan(&userExists)
	if err != nil || !userExists {
		return errors.New("user not found")
	}

	return nil
}

func isUserExists(dbConn *sql.DB, email string) (bool, error) {
	var userExists bool

	query := "SELECT EXISTS(SELECT 1 FROM user WHERE email = $1)"
	err := dbConn.QueryRow(query, email).Scan(&userExists)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return userExists, nil
}

func NewUserRepository(dbConn *sql.DB) IUserRepository {
	return &UserRepository{conn: dbConn}
}
