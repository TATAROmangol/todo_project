package postgres

import "fmt"

var(
	ErrHost = fmt.Errorf("not found env PG_HOST")
	ErrPort = fmt.Errorf("no found env PG_PORT")
	ErrUser = fmt.Errorf("no found env PG_USER")
	ErrPassword = fmt.Errorf("no found env PG_PASSWORD")
	ErrDBName = fmt.Errorf("no found env PG_DB_NAME")
	ErrSSL = fmt.Errorf("no found env PG_SSL")
)