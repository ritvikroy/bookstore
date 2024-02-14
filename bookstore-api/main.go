package main

import (
	"bookstore-api/db"
	"bookstore-api/router"
	"context"
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	dbConnection, dbError := db.DbConnection()
	if dbError != nil {
		fmt.Println("dbError is", dbError)
	}
	err := db.Migrations(context.TODO())
	if err != nil {
		fmt.Errorf("error occured %s", err)
	}
	appEngine := gin.Default()
	appEngine.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
	router.RegisterAllRoutes(dbConnection, appEngine)
	appEngine.Run("localhost:8080")

}
