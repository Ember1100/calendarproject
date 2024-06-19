package model

import "time"

type SysCalendar struct {
	ID uint `json:"id"`

	WarnDate time.Time `json:"warn_date"` //提醒日期

	WarnContext string `json:"warn_context" gorm:"type:text;not null"` //提醒内容

	CreateID uint `json:"create_id"` //创建人id

	CreatedAt time.Time `json:"create_at"`

	UpdatedAt time.Time `json:"update_at"`

	Status uint8 `json:"status" gorm:"type:tinyint;default:0"` //提醒发送状态

	SendType string `json:"send_type" gorm:"type:varchar(10)"` //接收信息方式 phone email
}

type SysCalendarSendRecord struct {
	ID uint `json:"id"`

	CreatedAt time.Time `json:"create_at"`

	UpdatedAt time.Time `json:"update_at"`

	SendAt time.Time `json:"send_at"`

	SysCalendarID uint `json:"sys_calendar_id"` //SysCalendar提醒表id
}
