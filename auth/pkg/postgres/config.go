package postgres

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSL string
}

func LoadConfig() (Config, error){
	host, exist := os.LookupEnv("PG_HOST")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_HOST")
	}

	sPort, exist := os.LookupEnv("PG_PORT")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_PORT")
	}
	port, err := strconv.Atoi(sPort)
	if err != nil{
		return Config{}, fmt.Errorf("failed load PG_PORT: %v", err)
	}

	user, exist := os.LookupEnv("PG_USER")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_USER")
	}

	password, exist := os.LookupEnv("PG_PASSWORD")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_PASSWORD")
	}

	dbName, exist := os.LookupEnv("PG_DB_NAME")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_DB_NAME")
	}

	ssl, exist := os.LookupEnv("PG_SSL")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_SSL")
	}

	return Config{
		host,
		port,
		user,
		password,
		dbName,
		ssl,
	}, nil
}

func (cfg Config) GetConnectPath() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSL)
}
