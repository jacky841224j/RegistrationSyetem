package main

import (
	routes "gotest/router"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
