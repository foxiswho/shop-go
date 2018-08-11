package models

type GoodsStatistics struct {
	Id      int     `json:"id" xorm:"not null pk INT(11)"`
	Saless  float64 `json:"saless" xorm:"not null default 0.0000 comment('销量') DECIMAL(28,4)"`
	Reading float64 `json:"reading" xorm:"not null default 0.0000 comment('访问数') DECIMAL(28,4)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewGoodsStatistics() *GoodsStatistics {
	return new(GoodsStatistics)
}
