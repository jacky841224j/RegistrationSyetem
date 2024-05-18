package services

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

type SchedulerService struct {
	cron *cron.Cron
}

func NewSchedulerService() *SchedulerService {
	return &SchedulerService{
		cron: cron.New(),
	}
}

func (s *SchedulerService) Start() {
	s.cron.AddFunc("CRON_TZ=Asia/Taipei 00 12 * * *", ReserveTask) // 每天12:30執行一次
	s.cron.Start()
}

func (s *SchedulerService) Stop() {
	s.cron.Stop()
}

// 預約通知排程
func ReserveTask() {
	DBOpen()
	//查詢明天的預約
	DB.Query("SELECT * FROM Calendar WHERE Date = ?", time.Now().AddDate(0, 0, 1))
	//通知明天有預約的客戶

	fmt.Println("Task executed at", time.Now())
	DBClose()
}
