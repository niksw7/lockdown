package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"io"
	"io/ioutil"
	"lockdown/handler"
	"lockdown/repository"
	"log"
	"os"
)

func main() {

	//get user and password
	username := os.Getenv("ADMIN")
	password := os.Getenv("PASSWORD")
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
	r.Use(RequestLoggerMiddleware())
	repo := repository.BuntDbRepo{DB: db}
	openEndpoints := r.Group("/open")
	openEndpoints.GET("/health", handler.HealthChecker())
	openEndpoints.POST("/register-user-details", handler.UserDetailsRegistrar(repo))

	authenticatedEndpoints := r.Group("/secure/", gin.BasicAuth(gin.Accounts{
		username: password,
	}))
	authenticatedEndpoints.GET("/read-user-details", handler.UserDetailsReader())
	authenticatedEndpoints.GET("/download-csv", handler.CsvDownloader(repo))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		log.Println("headers ", c.Request.Header)
		log.Println("Request Body :", string(body))
		c.Next()
	}
}
