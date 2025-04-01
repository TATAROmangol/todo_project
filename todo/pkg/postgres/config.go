package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSL      string
}

func MustLoadConfig() Config {
	host, exist := os.LookupEnv("POSTGRES_HOST")
	if !exist {
		log.Fatal("no found env POSTGRES_HOST")
	}

	sPort, exist := os.LookupEnv("POSTGRES_PORT")
	if !exist {
		log.Fatal("no found env POSTGRES_PORT")
	}
	port, err := strconv.Atoi(sPort)
	if err != nil {
		log.Fatal("invalid env POSTGRES_PORT")
	}

	user, exist := os.LookupEnv("POSTGRES_USER")
	if !exist {
		log.Fatal("no found env POSTGRES_USER")
	}

	password, exist := os.LookupEnv("POSTGRES_PASSWORD")
	if !exist {
		log.Fatal("no found env POSTGRES_PASSWORD")
	}

	dbName, exist := os.LookupEnv("POSTGRES_DB")
	if !exist {
		log.Fatal("no found env POSTGRES_DB")
	}

	ssl, exist := os.LookupEnv("POSTGRES_SSL")
	if !exist {
		log.Fatal("no found env POSTGRES_SSL")
	}

	return Config{
		Host: host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSL:      ssl,
	}
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSL,
	)
}
