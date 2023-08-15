package main

import (
	"main/internal/handlers"
	"main/model/handler"
	"net/http"
)

func main() {

	handler.InitRessources()

	http.HandleFunc("/", handlers.BaseHandler)

	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/loginConnection", handlers.LoginConnection)
	http.HandleFunc("/loginList", handlers.LoginListPage)

	http.HandleFunc("/register", handlers.RegisterHanlder)
	http.HandleFunc("/createAccount", handlers.RegisterNewAccount)

	http.ListenAndServe(":8080", nil)
}
