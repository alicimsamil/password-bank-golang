package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	dbConfig "password-bank-golang/config"
)

func GetDBConn() (*sql.DB, error) {
	config := dbConfig.GetDbConfig()
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.UserName, config.Password, config.DbName, config.SslMode)

	conn, err := sql.Open(config.DriverName, connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
