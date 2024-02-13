package db

import (
	"context"

	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitDb(ctxt context.Context) (*sql.Conn, error) {

	dbConnection, dbError := DbConnection(ctxt)
	if dbError != nil {
		return nil, dbError
	}
	dbConection, dbError := dbConnection.Conn(ctxt)
	if dbError != nil {
		fmt.Println("Db is not able to connect", dbError)
		return nil, dbError
	}
	return dbConection, nil
}

func DbConnection(ctxt context.Context) (*sql.DB, error) {
	dbOpen, dbOpenErr := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/bookstore?sslmode=disable")
	if dbOpenErr != nil {
		fmt.Println("dbOpenErr is", dbOpenErr)
		return nil, dbOpenErr
	}
	return dbOpen, nil
}

func Migrations(ctxt context.Context) error {
	dbConnection, dbError := DbConnection(ctxt)
	if dbError != nil {
		return dbError
	}
	driver, driverErr := postgres.WithInstance(dbConnection, &postgres.Config{})
	if driverErr != nil {
		fmt.Println("driver error is", driverErr)
		return driverErr
	}

	migration, migrationErr := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if migrationErr != nil {
		fmt.Println("Migration Error is", migrationErr.Error())
		return migrationErr
	}
	migrationUpErr := migration.Up()
	if migrationUpErr != nil {
		fmt.Println("migrationUpErr  is ", migrationUpErr)
		return migrationUpErr
	}
	return nil
}
