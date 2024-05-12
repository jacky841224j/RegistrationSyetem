package Controller

import (
	"RegistrationSyetem/Dto"
	"RegistrationSyetem/Services"
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

	var user Dto.User
	//
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = Services.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := Dto.ApiResponse{ResultCode: "200", ResultMessage: "會員新增成功"}
	Services.ResponseWithJson(w, http.StatusOK, response)
}
