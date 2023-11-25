package dbrepo

import (
	"database/sql"
	"udemyCourse1/internal/config"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) *postgresDBRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}