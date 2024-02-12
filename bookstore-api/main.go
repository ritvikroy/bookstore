package main

import (
	"bookstore-api/db"
	"bookstore-api/router"
	"context"
	"fmt"
)

func main() {
	connectionString := db.BuildDbDescriptiveConnString()

	db, err := db.InitDb(context.TODO(), connectionString)
	if err != nil {
		fmt.Errorf("error occured %s", err)
		return
	}
	appEngine := router.RegisterAllRoutes(db)
	appEngine.Run("localhost:8080")

}
