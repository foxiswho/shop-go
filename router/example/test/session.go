package test

import (
	"fmt"
	"github.com/foxiswho/shop-go/module/context"
)

//
type Session struct {

}

func NewSession() *Session{
	return  new(Session)
}

func (x *Session) IndexHandler(c *context.BaseContext) error {
	c.Session().Set("SSSSSSS","asldfjlksajdflkasjdflkjd")

	test:=c.Session().Get("SSSSSSS")
	fmt.Println("session_type=》SSSSSSS",test)

	c.Set("tmpl", "example/test/session_type")
	c.Set("data", map[string]interface{}{
		"title": "测试 COOIE",
		"test":  test,
	})

	return nil
}