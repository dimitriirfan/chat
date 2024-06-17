package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"github.com/caarlos0/env/v11"
)

type DependenciesConfig struct {
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBHost string `env:"DB_HOST"`
	DBName string `env:"DB_NAME"`
}

type Dependencies struct {
	DB *sql.DB
}

func InitializeAllDependencies() Dependencies {
	cfg := DependenciesConfig{}
	env.Parse(&cfg)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBName,
	))

	if err != nil {
		panic(err)
	}
	return Dependencies{
		DB: db,
	}
}
