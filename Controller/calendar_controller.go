package controller

import (
	"RegistrationSyetem/dto"
	"io"
	"net/http"
)

func CreateCalendar(w http.ResponseWriter, r *http.Request){ 
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}