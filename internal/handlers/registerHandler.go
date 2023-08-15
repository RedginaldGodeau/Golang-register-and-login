package handlers

import (
	"fmt"
	"main/model"
	"net/http"
	"os"
	"time"
)

func RegisterHanlder(w http.ResponseWriter, req *http.Request) {
	page := model.Page{Path: "register.html", Data: nil}
	page.ShowTemplate(w)
}

func RegisterNewAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/register", http.StatusAccepted)
		return
	}

	emailForm := req.FormValue("email")
	usernameForm := req.FormValue("username")
	passwordForm := req.FormValue("password")
	createOn := time.Now().UTC()

	database := model.Database{Driver: "postgres", Source: os.Getenv("POSTGRESQL_URL")}
	database.Connection()

	sql := "INSERT INTO account(email, username, password, create_on) VALUES($1, $2, $3, $4)"
	_, err := database.Db.Exec(sql, emailForm, usernameForm, passwordForm, createOn)
	defer database.Close()

	if err != nil {
		fmt.Fprintln(w, "[Create Bug] ", err)
	}

	fmt.Fprintln(w, "Account Create")
}
