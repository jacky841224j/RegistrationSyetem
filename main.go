package main

import (
	Routes "RegistrationSyetem/route"
	"net/http"
)

func main() {
	router := Routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
