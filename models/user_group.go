package models

type UserGroup struct {
	Id          int    `json:"id" xorm:"not null pk autoincr comment('会员用户组ID') INT(10)"`
	Name        string `json:"name" xorm:"default 'NULL' comment('名称') CHAR(30)"`
	Discount    int    `json:"discount" xorm:"not null default 0 comment('折扣率') INT(11)"`
	IsShowPrice int    `json:"is_show_price" xorm:"not null default 1 comment('是否显示价格1是0否') TINYINT(1)"`
	Remark      string `json:"remark" xorm:"default 'NULL' comment('备注') VARCHAR(255)"`
	Sort        int    `json:"sort" xorm:"not null default 0 comment('排序') index INT(5)"`
	IsDel       int    `json:"is_del" xorm:"not null default 0 comment('是否删除1是0否') index TINYINT(1)"`
	Mark        string `json:"mark" xorm:"default 'NULL' comment('标志') CHAR(15)"`
	Qq          string `json:"qq" xorm:"default 'NULL' comment('客服') VARCHAR(15)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewUserGroup() *UserGroup {
	return new(UserGroup)
}
