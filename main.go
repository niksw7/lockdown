package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"lockdown/controllers"
	"log"
)

func main() {
	r := gin.Default()
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	r.Use(func(context *gin.Context) {
		context.Set("db", db)
	})
	openEndpoints := r.Group("/open")
	openEndpoints.GET("/health", controllers.Health)
	openEndpoints.POST("/register-user-details", controllers.RegisterUserDetails)

	//authenticatedEndpoints := r.Group("/secure/")
	r.GET("/read-data", controllers.ReadUserRegisteredDetails)
	r.GET("/download-csv",controllers.DownloadCsv)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
