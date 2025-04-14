package migrator

import "fmt"

var(
	ErrHost = fmt.Errorf("not found env MIGRATE_HOST")
	ErrPort = fmt.Errorf("no found env MIGRATE_PORT")
	ErrUser = fmt.Errorf("no found env MIGRATE_USER")
	ErrPassword = fmt.Errorf("no found env MIGRATE_PASSWORD")
	ErrDBName = fmt.Errorf("no found env MIGRATE_DB_NAME")
	ErrSSL = fmt.Errorf("no found env MIGRATE_SSL")
	ErrSchema = fmt.Errorf("no found env MIGRATE_SCHEMA")
)