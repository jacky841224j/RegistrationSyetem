package controller

import (
	"RegistrationSyetem/dto"
	"RegistrationSyetem/services"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//讀取參數
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user dto.User
	//
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = services.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := dto.ApiResponse{ResultCode: "200", ResultMessage: "會員新增成功"}
	services.ResponseWithJson(w, http.StatusOK, response)
}
