package session_models

type User struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Mobile   string `json:"mobile" xorm:"default 'NULL' index CHAR(11)"`
	Username string `json:"username" xorm:"default 'NULL' comment('用户名') index CHAR(30)"`
	Mail     string `json:"mail" xorm:"default 'NULL' comment('邮箱') index CHAR(32)"`
	GroupId  int    `json:"group_id" xorm:"not null default 410 comment('用户组ID') index INT(11)"`
	Name     string `json:"name" xorm:"default 'NULL' comment('店铺名称') VARCHAR(100)"`
}
