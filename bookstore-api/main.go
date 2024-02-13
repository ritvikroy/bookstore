package main

import (
	"bookstore-api/db"
	"context"
	"fmt"
	"net/http"
	"bookstore-api/router"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// dbConnection, dbError := db.DbConnection(context.TODO())
	// if dbError != nil {
	// 	fmt.Println("dbError is", dbError)
	// }
	err := db.Migrations(context.TODO())
	if err != nil {
		fmt.Println("Error is ", err)
	}

	// db, err := db.InitDb(context.TODO(), connectionString)
	if err != nil {
		fmt.Errorf("error occured %s", err)
	}

	appEngine := router.RegisterAllRoutes()
	appEngine.StaticFS("/", http.Dir("../frontend/build"))

	appEngine.Run("localhost:8080")

}
