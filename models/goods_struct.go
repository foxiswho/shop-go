package models

type GoodsStruct struct {
	Goods      *Goods
	GoodsPrice *GoodsPrice
}

//初始化
func NewGoodsStruct() *GoodsStruct {
	return new(GoodsStruct)
}
