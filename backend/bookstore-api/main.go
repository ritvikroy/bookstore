package main

import (
	"bookstore-api/router"
)

func main() {
	//db, err := db.InitDB()
	// if err != nil {
	// 	fmt.Errorf("error occured %s", err)
	// 	return
	// }
	appEngine := router.RegisterAllRoutes()
	appEngine.Run("localhost:8080")
}
