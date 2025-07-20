package main

import (
	"log"
	"io"
	"net/http"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func router01() http.Handler {
	
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET("/hello", func(context *gin.Context) {
		log.Printf("Sent request to API2")
		message := getMessage()
		context.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": message,
			},
		)
		log.Printf("Get response from API2")
	})

	return engine
}

func getMessage() string {
	res, err := http.Get("http://api2:8081/message")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	type Message struct { Message string }

	dec := json.NewDecoder(strings.NewReader(string(body)))
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			return "error"
		} else if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", m.Message)
		return m.Message
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server01.ListenAndServe()
}