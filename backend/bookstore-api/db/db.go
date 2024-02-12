package db

import (
	"fmt"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

func buildDbDescriptiveConnString() string {
	dbPasswordKey := "password"
	dbSourceName := fmt.Sprintf("user=%s password=\"%s\" connectString=\"(DESCRIPTION=(CONNECT_DATA=(SERVICE_NAME=%s))(ADDRESS=(PROTOCOL=tcp)(HOST=%s)(PORT=%v)))\"",
		"bookstorej",
		dbPasswordKey,
		"XEPDB1",
		"localhost",
		1521)
	fmt.Println("dbSourceName: ", dbSourceName)
	return dbSourceName
}

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", buildDbDescriptiveConnString())
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Db connected Successfully")
	return db, nil
}
