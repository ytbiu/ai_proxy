package router

import (
	"ai_proxy/config"
	"ai_proxy/controller"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(r *gin.Engine) {
	for _, path := range config.ConfigInfo.AITaskExecutorNodeProxyAPIPath {
		r.Any(path, controller.Proxy)
	}

	// test proxy
	r.Any("/b/v1/models", func(c *gin.Context) {
		type rr struct {
			Name string
			Age  int
		}
		var a rr
		if err := c.Bind(&a); err != nil {
			panic("err")
		}
		logrus.Infof("%+v", a)

		for k, v := range c.Request.Header {
			logrus.Infof("k : %s, v: %s", k, v)
		}

		b, _ := json.Marshal(a)
		c.Writer.Write(b)
	})
}
