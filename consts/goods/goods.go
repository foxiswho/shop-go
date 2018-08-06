package goods

const (
	//商品类型
	TYPE_Default  = 10001 //商品
	TYPE_Integral = 10002 //积分
	//标志类别
	Mark_Default  = 10101 //默认商品 产品-仓库-供应商
	Mark_Combined = 10102 //组合商品 商品ID-10002
	//商品发货属性类别
	Type_Delivery      = 10201 //现货
	Type_Bonded_area   = 10202 //保税区
	Type_Direct_mail   = 10203 //直邮
	Type_General_trade = 10204 //一般贸易
)

//////////////////////////////////////////////////////////
//goods_price

const (
	//价格类别
	Price_Type_Default  = 21001 //默认
	Price_Type_Discount = 21002 //优惠
)
