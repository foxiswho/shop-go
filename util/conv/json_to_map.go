package conv

import "github.com/foxiswho/shop-go/util/json"

func JsonToMap(str string, m map[string]interface{}) error {
	return json.Unmarshal([]byte(str), &m)
}
