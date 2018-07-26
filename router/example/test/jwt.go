package test

import "github.com/foxiswho/shop-go/module/context"

func JwtTester2Handler(c *context.BaseContext) error {
	c.Set("tmpl", "example/test/jwt_tester")
	c.Set("data", map[string]interface{}{
		"title": "JWT 接口测试",
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiaWQiOiIxIiwibmFtZSI6IkhvYm8ifQ.YUzBykoELyKoQWaugkVNf3d09HBhICBJoOcWQKnveRQ",
	})

	return nil
}