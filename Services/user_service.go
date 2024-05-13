package services

import (
	"RegistrationSyetem/dto"
	"fmt"
)

func CreateUser(user dto.User) error {

	DBOpen()

	stmt, err := DB.Prepare("INSERT INTO member (name,birth) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Birth)
	if err != nil {
		return err
	}

	fmt.Println("User新增成功", user)
	return nil
}
