package task

import (
	"fmt"
	"time"

	"github.com/calendarproject/common"
	"github.com/calendarproject/model"
	"gorm.io/gorm"
)

type SysTask struct {
	DB *gorm.DB
}

func NewSysTask() SysTask {
	db := common.GetDb()
	return SysTask{DB: db}
}

func (s SysTask) SendMessTask() {
	var posts []model.SysCalendar
	nowTime := time.Now()
	s.DB.Where("status = ? and warn_date <= ?  ", 0, nowTime).Find(&posts)
	if len(posts) == 0 {
		return
	}
	for i := 0; i < len(posts); i++ {
		value := posts[i]
		if value.SendType == "" || len(value.SendType) == 0 {
			continue
		} else if value.SendType == "email" {
			SendByEmail(value)
		} else if value.SendType == "phone" {
			SendByPhone(value)
		}
		posts[i].Status = 1
	}
	var record []model.SysCalendarSendRecord
	//更新
	for _, value := range posts {
		if value.Status == 1 {
			record = append(record, model.SysCalendarSendRecord{SysCalendarID: value.ID, SendAt: nowTime})
			if err := s.DB.Model(&model.SysCalendar{}).Where("id = ?", value.ID).Update("status", 1).Error; err != nil {
				fmt.Println(value.ID, "更新失败")
				fmt.Println(err)
			}
		}
	}

	//记录发送
	s.SendRecord(record, nowTime)
}

// 邮箱发送
func SendByEmail(calendar model.SysCalendar) {

}

// 手机发送
func SendByPhone(calendar model.SysCalendar) {

}

func (s SysTask) SendRecord(record []model.SysCalendarSendRecord, nowTime time.Time) {
	if len(record) == 0 {
		return
	}
	s.DB.AutoMigrate(&model.SysCalendarSendRecord{})
	s.DB.Create(&record)
}
