package migrator

import (
	"fmt"
	"os"
)

type Config struct{
	address string
}

func Load() (Config, error){
	user, exist := os.LookupEnv("MIGRATE_USER")
	if !exist{
		return Config{}, ErrUser
	}

	password, exist := os.LookupEnv("MIGRATE_PASSWORD")
	if !exist{
		return Config{}, ErrPassword
	}
	
	host, exist := os.LookupEnv("MIGRATE_HOST")
	if !exist{
		return Config{}, ErrHost
	}

	port, exist := os.LookupEnv("MIGRATE_PORT")
	if !exist{
		return Config{}, ErrPort
	}

	dbName, exist := os.LookupEnv("MIGRATE_DB_NAME")
	if !exist{
		return Config{}, ErrDBName
	}

	ssl, exist := os.LookupEnv("MIGRATE_SSL")
	if !exist{
		return Config{}, ErrSSL
	}

	schema, exist := os.LookupEnv("MIGRATE_SCHEMA")
	if !exist{
		return Config{}, ErrSchema
	}

	address := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%v&x-migrations-table=%s",
		user, password, host, port, dbName, ssl, schema,
	)
	return Config{address}, nil
}