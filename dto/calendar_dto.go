package dto

import (
	"time"
)

type Calendar struct {
	Id      int       //pk
	Date    time.Time //日期
	Time    time.Time //時間
	UserID  int       //fk
	Content string    //診療內容
}
