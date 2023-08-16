package main

import (
	"fmt"
	"main/src/model/handler"
	"main/src/model/router"
	"net/http"
)

func main() {
	handler.InitRessources()
	router.RouteMaker()

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
