package handlers

import (
	"fmt"
	"main/src/domain/entity/account"
	"main/src/model"
	"main/src/model/token"
	"net/http"
	"time"
)

func (h *Handler) LoginPage(w http.ResponseWriter, req *http.Request) {
	page := model.Page{Path: "login.html", Data: nil}
	page.ShowTemplate(w)
}

func (h *Handler) LoginListPage(w http.ResponseWriter, req *http.Request) {
	accoutsList := account.GetAll()
	if accoutsList == nil {
		fmt.Fprintln(w, "Database Error")
		return
	}

	fmt.Fprintln(w, "List of accounts: ")
	for _, account := range *accoutsList {
		fmt.Fprintln(w, account.Id, account.Username, account.Email, account.Password)
	}
}

func (h *Handler) LoginConnection(w http.ResponseWriter, req *http.Request) {

	emailForm := req.FormValue("email")
	passwordForm := req.FormValue("password")

	account := account.GetByEmail(emailForm)
	if account == nil {
		fmt.Fprintln(w, "Email not found")
		return
	}

	if account.Password == passwordForm {

		to, err := token.New(account.Username, "login")
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		cookie := http.Cookie{}
		cookie.Name = "accessToken"
		cookie.Value = *to
		cookie.Expires = time.Now().Add(10 * time.Minute)
		cookie.Secure = false
		http.SetCookie(w, &cookie)

		fmt.Fprintln(w, "Hello you :)")
		return
	}

	fmt.Fprintln(w, "Incorrect password")
}
