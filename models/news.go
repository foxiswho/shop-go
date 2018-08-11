package models

import (
	"time"
)

type News struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Aid         int       `json:"aid" xorm:"not null default 0 comment('管理员AID') INT(11)"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('是否删除1是0否') index(is_del) TINYINT(1)"`
	IsOpen      int       `json:"is_open" xorm:"not null default 1 comment('启用1是0否') index(is_del) TINYINT(1)"`
	Status      int       `json:"status" xorm:"not null default 0 comment('状态') index(is_del) INT(11)"`
	GmtSystem   time.Time `json:"gmt_system" xorm:"default 'NULL' comment('创建时间,系统时间不可修改') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'current_timestamp()' comment('添加时间,可修改') TIMESTAMP"`
	Title       string    `json:"title" xorm:"default 'NULL' comment('标题') VARCHAR(255)"`
	Author      string    `json:"author" xorm:"default 'NULL' comment('作者') VARCHAR(255)"`
	Url         string    `json:"url" xorm:"default 'NULL' comment('网址') VARCHAR(255)"`
	UrlSource   string    `json:"url_source" xorm:"default 'NULL' comment('来源地址(转载)') VARCHAR(255)"`
	UrlRewrite  string    `json:"url_rewrite" xorm:"default 'NULL' comment('自定义伪静态Url') index CHAR(150)"`
	Description string    `json:"description" xorm:"default 'NULL' comment('摘要') VARCHAR(255)"`
	Content     string    `json:"content" xorm:"default 'NULL' comment('内容') TEXT"`
	Type        int       `json:"type" xorm:"not null default 0 comment('类型0文章10001博客栏目') index INT(11)"`
	ModuleId    int       `json:"module_id" xorm:"not null default 0 comment('模块10019技术10018生活') index INT(10)"`
	SourceId    int       `json:"source_id" xorm:"not null default 0 comment('来源:后台，接口，其他') index INT(11)"`
	TypeId      int       `json:"type_id" xorm:"not null default 0 comment('类别ID，原创，转载，翻译') index(is_del) INT(11)"`
	CatId       int       `json:"cat_id" xorm:"not null default 0 comment('分类ID，栏目') index(is_del) INT(11)"`
	Tag         string    `json:"tag" xorm:"default 'NULL' comment('标签') VARCHAR(255)"`
	Thumb       string    `json:"thumb" xorm:"default 'NULL' comment('缩略图') VARCHAR(255)"`
	IsRelevant  int       `json:"is_relevant" xorm:"not null default 0 comment('相关文章1是0否') TINYINT(1)"`
	IsJump      int       `json:"is_jump" xorm:"not null default 0 comment('跳转1是0否') TINYINT(1)"`
	IsComment   int       `json:"is_comment" xorm:"not null default 1 comment('允许评论1是0否') TINYINT(1)"`
	IsRead      int       `json:"is_read" xorm:"not null default 10014 comment('是否阅读10014未看10015在看10016已看') INT(11)"`
	Sort        int       `json:"sort" xorm:"not null default 0 comment('排序') index(is_del) INT(11)"`
	Remark      string    `json:"remark" xorm:"default 'NULL' comment('备注') VARCHAR(255)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewNews() *News {
	return new(News)
}
