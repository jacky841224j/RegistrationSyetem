package dto

import (
	"time"
)

type GetCalendar struct {
	Id        *int       //pk
	StartDate *time.Time //開始日期
	EndDate   *time.Time //結束日期
	StartTime *time.Time //開始時間
	EndTime   *time.Time //結束時間
	UserID    *int       //fk
	Content   *string    //診療內容
}
