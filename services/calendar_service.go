package services

import (
	"RegistrationSyetem/dto"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// 新增預約
func CreateCalendar(dto dto.Calendar) (bool, error) {

	//連接資料庫
	DBOpen()

	//判斷預約時段與行事曆時間是否有衝突
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM calendar WHERE date = ? AND time BETWEEN ? AND ?", dto.Date, dto.Time.Add(-30*time.Minute), dto.Time.Add(30*time.Minute)).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, errors.New("與預約時間衝突，預約失敗")
	}

	//新增預約
	_, err = DB.Exec("INSERT INTO member (date,time,userid,content) VALUES (?,?,?,?)", dto.Date, dto.Time, dto.UserID, dto.Content)
	if err != nil {
		return false, err
	}

	fmt.Println("Calendar新增成功", dto)
	return true, nil
}

// 取得預約
func GetCalendar(getCalendar dto.GetCalendar) []dto.Calendar {

	query := "SELECT * FROM calendar"
	var conditions []string
	var args []interface{}

	if getCalendar.UserID != nil {
		conditions = append(conditions, "userid = ?")
		args = append(args, *getCalendar.UserID)
	}

	if getCalendar.Id != nil {
		conditions = append(conditions, "id = ?")
		args = append(args, *getCalendar.Id)
	}

	if getCalendar.StartDate != nil && getCalendar.EndDate != nil {
		conditions = append(conditions, "date >= ? and date <= ?")
		args = append(args, *getCalendar.StartDate)
		args = append(args, *getCalendar.EndDate)
	}

	if getCalendar.StartTime != nil && getCalendar.EndTime != nil {
		conditions = append(conditions, "time >= ? and time <= ?")
		args = append(args, *getCalendar.StartTime)
		args = append(args, *getCalendar.EndTime)
	}

	if len(conditions) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditions, " AND "))
	}

	rows, err := DB.Query(query, args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 讀取查詢結果
	var results []dto.Calendar
	for rows.Next() {
		var calendar dto.Calendar
		if err := rows.Scan(&calendar.Id, &calendar.Date, &calendar.Time, &calendar.UserID, &calendar.RemedyItem, &calendar.Content, &calendar.Remark); err != nil {
			log.Fatal(err)
		}
		results = append(results, calendar)
	}

	// 檢查是否有查詢錯誤
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}
