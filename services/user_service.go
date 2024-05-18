package services

import (
	"RegistrationSyetem/dto"
	"fmt"
)

func CreateUser(dto dto.User) error {

	//連接資料庫
	DBOpen()

	stmt, err := DB.Prepare("INSERT INTO member (name,birth) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dto.Name, dto.Birth)
	if err != nil {
		return err
	}

	fmt.Println("User新增成功", dto)
	return nil
}
