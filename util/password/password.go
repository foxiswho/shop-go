package password

import "github.com/foxiswho/shop-go/util/crypt"

//密码加密
func SaltMake(str, salt string) string {
	return crypt.Md5(crypt.Sha256(str+salt) + salt)
}
