package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//connectionString := db.BuildDbDescriptiveConnString()
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)

	//db, err := db.InitDb(context.TODO(), connectionString)
	if err != nil {
		fmt.Errorf("error occured %s", err)
		return
	}
	//appEngine := router.RegisterAllRoutes(db)
	appEngine := gin.Default()
	appEngine.StaticFS("/", http.Dir("./static"))

	appEngine.Run("localhost:8080")

}
