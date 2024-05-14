package services

import (
	"RegistrationSyetem/dto"
	"fmt"
)

func CreateCalendar(dto dto.Calendar) error {
	//連接資料庫
	DBOpen()

	stmt, err := DB.Prepare("INSERT INTO member (date,calendar) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dto.Date, dto.Calendar)
	if err != nil {
		return err
	}

	fmt.Println("User新增成功", dto)
	return nil
}
