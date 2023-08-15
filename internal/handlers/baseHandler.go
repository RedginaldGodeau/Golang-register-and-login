package handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"main/model/token"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, req *http.Request) {

	accessToken, err := req.Cookie("accessToken")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusAccepted)
		return
	}

	t, err := token.IsValid(accessToken.Value, "login")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusAccepted)
		return
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Redirect(w, req, "/login", http.StatusAccepted)
		return
	}

	username := claims["values"].(string)

	fmt.Fprintln(w, "Welcome Home ", username)
}
