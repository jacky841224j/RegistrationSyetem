package controller

import (
	"RegistrationSyetem/dto"
	"RegistrationSyetem/services"
	"encoding/json"
	"io"
	"net/http"
)

func CreateCalendar(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var calendar dto.Calendar

	err = json.Unmarshal(b, &calendar)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := dto.ApiResponse{ResultCode: "200", ResultMessage: "會員新增成功"}
	services.ResponseWithJson(w, http.StatusOK, response)
}
