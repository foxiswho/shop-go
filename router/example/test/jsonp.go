package test

import (
	"github.com/foxiswho/shop-go/module/context"
	"time"
	"math/rand"
	"net/http"
)

type Json struct {
}

type A struct {
	F  string `json:"f,filter:*"`
	F1 string `json:"f_1,filter:a1"`
	F2 string `json:"f_2,filter:a2"`

	B  B `json:"b,filter:*.*"`
	B1 B `json:"b_1,filter:*.b1"`
	B2 B `json:"b_2,filter:a2.b2"`
	B3 B `json:"b_3,filter:a1.*;a2.a2"`
}

type B struct {
	F  string `json:"f,filter:*"`
	F1 string `json:"f_1,filter:b1"`
	F2 string `json:"f_2,filter:b2"`

	C  C `json:"c,filter:*.*"`
	C1 C `json:"c_1,filter:b1.*"`
	C2 C `json:"c_3,filter:*.b1"`
}

type C struct {
	F  string `json:"f,filter:*"`
	F1 string `json:"f_1,filter:b1"`
	F2 string `json:"f_2,filter:b2"`
}

func JsonpIndexHandler(c *context.BaseContext) error {
	c.Set("tmpl", "example/test/jsonp")
	c.Set("data", map[string]interface{}{
		"title": "jsonp",
	})
	return nil
}

func JsonpIndexShowHandler(c *context.BaseContext) error {
	callback := c.QueryParam("callback")
	var content struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}
	content.Response = "Sent via JSONP"
	content.Timestamp = time.Now().UTC()
	content.Random = rand.Intn(1000)

	return c.JSONP(http.StatusOK, callback, &content)
}