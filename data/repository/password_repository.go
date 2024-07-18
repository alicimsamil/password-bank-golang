package repository

import "database/sql"

type IPasswordRepository interface {
}

type PasswordRepository struct {
	conn *sql.Conn
}

func NewPasswordRepository(dbConn *sql.Conn) IPasswordRepository {
	return &PasswordRepository{conn: dbConn}
}
