package dto

import (
	"time"
)

type Calendar struct {
	Id         int       //pk
	Date       time.Time //日期
	Time       time.Time //時間
	UserID     int       //fk
	RemedyItem string    //治療項目
	Content    string    //診療內容
	Remark     string    //備註
}
