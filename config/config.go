package config

// SecretKey JWT secret
var SecretKey = []byte("your-secret-key")

type DBConfig struct {
	Host       string
	Port       int16
	UserName   string
	Password   string
	DbName     string
	SslMode    string
	DriverName string
}

func GetDbConfig() DBConfig {
	return DBConfig{
		Host:       "localhost",
		Port:       5432,
		UserName:   "postgres",
		Password:   "PASSWORD",
		DbName:     "password_bank",
		SslMode:    "disable",
		DriverName: "postgres",
	}
}
