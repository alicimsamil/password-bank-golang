package repository

import "database/sql"

type IUserRepository interface {
}

type UserRepository struct {
	conn *sql.Conn
}

func NewUserRepository(dbConn *sql.Conn) IUserRepository {
	return &UserRepository{conn: dbConn}
}
