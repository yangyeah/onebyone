package onebyone

import (
	"encoding/json"
	"strings"

	"github.com/cdle/sillyGirl/core"
	"github.com/gin-gonic/gin"
)

type AutoGenerated struct {
	PtPin   string `json:"pt_pin"`
	Message string `json:"message"`
}

func init() {

	core.Server.POST("/onebyone/push", func(c *gin.Context) {
		data, _ := c.GetRawData()
		ag := &AutoGenerated{}
		json.Unmarshal(data, ag)
		ptPin := ag.PtPin
		message := ag.Message
		for _, tp := range []string{
			"qq", "tg", "wx",
		} {
			core.Bucket("pin" + strings.ToUpper(tp)).Foreach(func(k, v []byte) error {
				if string(k) == ptPin && ptPin != "" {
					if push, ok := core.Pushs[tp]; ok {
						push(string(v), message)
					}
				}
				return nil
			})
		}
		c.String(200, "ok")
	})
}
