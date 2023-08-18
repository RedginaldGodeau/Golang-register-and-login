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
	accountsList := account.GetAll()
	if accountsList == nil {
		fmt.Fprintln(w, "Database Error")
		return
	}

	fmt.Fprintln(w, "List of accounts: ")
	for _, acc := range *accountsList {
		fmt.Fprintln(w, acc.Id, acc.Username, acc.Email, acc.Password)
	}
}

func (h *Handler) LoginConnection(w http.ResponseWriter, req *http.Request) {

	emailForm := req.FormValue("email")
	passwordForm := req.FormValue("password")

	acc := account.GetByEmail(emailForm)
	if acc == nil {
		fmt.Fprintln(w, "Email not found")
		return
	}

	if acc.Password == passwordForm {

		to, err := token.New(acc.Username, "login")
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
