package main

import (
	"api/util"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load config from .env file
	config, err := util.Loadconfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect to DB using config values from .env file
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	fmt.Println("Connected to DB", db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := "8080"
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
