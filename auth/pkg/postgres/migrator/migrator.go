package migrator

import (
	"auth/pkg/postgres"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct{
	m *migrate.Migrate
}

func New(dirPath string, dbCfg postgres.Config) (*Migrator, error){
	dbUrl := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%v",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DBName, dbCfg.SSL,
	)

	m, err := migrate.New(dirPath, dbUrl)
	if err != nil{
		return nil, fmt.Errorf("failed create migrator, err: %v", err)
	}

	return &Migrator{m}, nil
}

func (mig *Migrator) Up() error {
	defer mig.m.Close()

	err := mig.m.Up()
	if err == nil || err == migrate.ErrNoChange{
		return nil
	}

	version, _, _ := mig.m.Version()
	vers := int(version) - 1
	if err := mig.m.Force(vers); err != nil{
		return fmt.Errorf("failed rollback migration: err=%v", err)
	}

	return fmt.Errorf("migrations are not applied: current version=%v, err=%v", vers, err)
}