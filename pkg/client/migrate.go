package client

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

func Migrate(pgx *pgxpool.Pool) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(pgx)

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	log.Println("Migration completed successfully")

	return nil
}
