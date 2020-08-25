package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", http.NotFound)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user", user)

	logrus.WithField("port", 8080).Info("Starting HTTP Server")
	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("Path", r.URL.Path).WithField("Method", r.Method).Info("call hello")
	w.Write([]byte("Hello World"))
}

func user(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("Path", r.URL.Path).WithField("Method", r.Method).Info("call user")
	switch r.Method {
	case "GET":
		user := User{1, "jayden"}
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(user)

	case "POST":
		var user User
		json.NewDecoder(r.Body).Decode(&user)

	default:

	}
}
