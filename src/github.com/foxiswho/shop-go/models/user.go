package models

import (
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty" xorm:"pk autoincr"`
	Nickname  string    `form:"nickname" json:"nickname,omitempty"`
	Password  string    `form:"password" json:"-"`
	Gender    int64     `json:"gender,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty" xorm:"'created_time'" `
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty" xorm:"'updated_time'"`
}
