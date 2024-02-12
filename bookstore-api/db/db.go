package db

import (
	"context"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

func BuildDbDescriptiveConnString() string {
	dbPasswordKey := "postgres"
	dbName := "bookshop-db"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", dbPasswordKey, dbName)
	fmt.Println("dbSourceName: ", psqlInfo)
	return psqlInfo
}

func InitDb(ctxt context.Context, dbSourceName string) (*sql.Conn, error) {
	dbClient, dbError := sql.Open("postgres", dbSourceName)
	if dbError != nil {
		fmt.Println("Db is not able to open", dbError)
		return nil, dbError
	}
	dbConection, dbError := dbClient.Conn(ctxt)
	if dbError != nil {
		fmt.Println("Db is not able to connect", dbError)
		return nil, dbError
	}
	return dbConection, nil
}
