package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"lockdown/handler"
	"lockdown/repository"
	"log"
)

func main() {
	r := gin.Default()
	db, err := buntdb.Open("data.db")
	db.CreateIndex("jsonIndex", "*", buntdb.IndexJSON("Id"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	r.Use(func(context *gin.Context) {
		context.Set("db", db)
	})
	repo := repository.BuntDbRepo{DB: db}
	openEndpoints := r.Group("/open")
	openEndpoints.GET("/health", handler.HealthChecker())
	openEndpoints.POST("/register-user-details", handler.UserDetailsRegistrar(repo))

	//authenticatedEndpoints := r.Group("/secure/")
	r.GET("/read-user-details", handler.UserDetailsReader())
	r.GET("/download-csv", handler.CsvDownloader(repo))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
