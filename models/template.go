package models

import (
	"time"
)

type Template struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('模板ID') INT(11)"`
	Name        string    `json:"name" xorm:"default 'NULL' comment('模板名称(中文)') VARCHAR(80)"`
	Mark        string    `json:"mark" xorm:"default 'NULL' comment('模板名称标志(英文)（调用时使用）') VARCHAR(80)"`
	Title       string    `json:"title" xorm:"default 'NULL' comment('邮件标题') VARCHAR(255)"`
	Type        int       `json:"type" xorm:"not null default 0 comment('模板类型1短信模板2邮箱模板') TINYINT(1)"`
	Use         int       `json:"use" xorm:"not null default 0 comment('用途') TINYINT(2)"`
	Content     string    `json:"content" xorm:"default 'NULL' TEXT"`
	Remark      string    `json:"remark" xorm:"default 'NULL' comment('备注') VARCHAR(1024)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'current_timestamp()' comment('创建时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`
	CodeNum     int       `json:"code_num" xorm:"not null default 0 comment('验证码位数') TINYINT(3)"`
	Aid         int       `json:"aid" xorm:"not null default 0 comment('添加人') INT(11)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewTemplate() *Template {
	return new(Template)
}
