package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/message", func(c *gin.Context) {
		log.Printf("Get request from API1")
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"Message": "Hello world",
				"AIP2 send": true,
			},
		)
		log.Printf("Send mesaage to API1")
	})

	return e
}

func main() {
	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02.ListenAndServe()
}