package api


import (
	"time"
	"math/rand"
	"net/http"
	"github.com/foxiswho/shop-go/module/context"
)


func JsonpHandler(c *context.BaseContext) error {
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