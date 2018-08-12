package number

import "github.com/foxiswho/echo-go/util/str"

//订单号
func OrderNoMake() string {
	return str.MakeYearDaysRand(12)
}
