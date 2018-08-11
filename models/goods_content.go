package models

type GoodsContent struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	SeoTitle       string `json:"seo_title" xorm:"default 'NULL' comment('seo标题') VARCHAR(50)"`
	SeoDescription string `json:"seo_description" xorm:"default 'NULL' comment('seo描述') VARCHAR(200)"`
	SeoKeyword     string `json:"seo_keyword" xorm:"default 'NULL' comment('seo关键字') VARCHAR(50)"`
	Content        string `json:"content" xorm:"default 'NULL' comment('内容') TEXT"`
	Remark         string `json:"remark" xorm:"default 'NULL' comment('备注紧供自己查看') VARCHAR(255)"`
	TitleOther     string `json:"title_other" xorm:"default 'NULL' comment('其他名称') VARCHAR(5000)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewGoodsContent() *GoodsContent {
	return new(GoodsContent)
}
