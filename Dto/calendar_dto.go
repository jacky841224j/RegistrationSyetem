package dto

import (
	"time"
)

type Calendar struct {
	Id       int
	Date     time.Time
	Calendar string
}
