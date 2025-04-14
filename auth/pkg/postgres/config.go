package postgres

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSL      string
}

func Load() (Config, error) {
	host, exist := os.LookupEnv("PG_HOST")
	if !exist {
		return Config{}, ErrHost
	}

	port, exist := os.LookupEnv("PG_PORT")
	if !exist {
		return Config{}, ErrPort
	}

	user, exist := os.LookupEnv("PG_USER")
	if !exist {
		return Config{}, ErrUser
	}

	password, exist := os.LookupEnv("PG_PASSWORD")
	if !exist {
		return Config{}, ErrPassword
	}

	dbName, exist := os.LookupEnv("PG_DB_NAME")
	if !exist {
		return Config{}, ErrDBName
	}

	ssl, exist := os.LookupEnv("PG_SSL")
	if !exist {
		return Config{}, ErrSSL
	}

	return Config{
		Host: host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSL:      ssl,
	}, nil
}

func (c *Config) GetConnString() string {
	return fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSL,
	)
}