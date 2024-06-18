package model

import "time"

type SysCalendar struct {
	ID uint `json:"id"`

	WarnDate time.Time `json:"warn_date"` //提醒日期

	WarnContext string `json:"warn_context" gorm:"type:text;not null"` //提醒内容

	CreateID uint `json:"create_id"` //创建人id

	CreatedAt time.Time `json:"create_at"`

	UpdatedAt time.Time `json:"update_at"`
}
