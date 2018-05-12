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
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty" xorm:"'created_time' default 'CURRENT_TIMESTAMP' TIMESTAMP" `
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty" xorm:"'updated_time'"`
}

//初始化
func NewUser() *User {
	return new(User)
}

//初始化列表
func (c *User) newMakeDataArr() []User {
	return make([]User, 0)
}
