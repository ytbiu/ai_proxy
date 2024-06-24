package controller

import (
	"ai_proxy/service/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Proxy(c *gin.Context) {
	path := c.Request.URL.Path
	l := logrus.WithField("path", path)
	if err := common.GinProxy(c); err != nil {
		l.Error("common.GinProxy err : ", err)
		return
	}
}
