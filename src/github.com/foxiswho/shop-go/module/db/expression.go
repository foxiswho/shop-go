package db

//查询条件
type QueryCondition struct {
	Field     string      //字段
	Operation string      //操作 =,!=,>,<,>=,<=,in,not in
	Condition interface{} //具体值
}

var expression map[int]QueryCondition

func NewMakeQueryCondition() map[int]QueryCondition {
	return make(map[int]QueryCondition)
}

//添加条件
func AddQueryCondition(field, qe string, condition interface{}) *QueryCondition {
	return &QueryCondition{field, qe, condition}
}
