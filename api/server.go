package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users = []User{}

type Router struct{}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if path == "/users" && req.Method == "GET" {
		response, _ := json.Marshal(Users)
		res.Write(response)
	} else if path == "/users" && req.Method == "POST" {
		var user User
		_ = json.NewDecoder(req.Body).Decode(&user)
		Users = append(Users, user)
	} else {
		http.NotFound(res, req)
	}
}

func RunService() {
	c := cors.AllowAll()

	handler := c.Handler(Router{})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
