package main

import (
	"bookstore-api/db"
	"bookstore-api/router"
	"context"
	"fmt"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	
	fmt.Println("logger to message")
	//_, dbError := db.DbConnection(context.TODO())
	//if dbError != nil {
	//	fmt.Println("dbError is", dbError)
	//}
	//err := db.Migrations(context.TODO())
	//if err != nil {
	//	fmt.Println("Error is ", err)
	//}
	//
	db, err := db.InitDb(context.TODO())
	if err != nil {
		fmt.Errorf("error occured %s", err)
	}
	
	fmt.Println("Init db in main.go: ", db)

	appEngine := router.RegisterAllRoutes()
	appEngine.StaticFS("/", http.Dir("../frontend/build"))

	appEngine.Run("localhost:8080")

}
