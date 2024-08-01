package controller

import (
	"ai_proxy/config"
	"ai_proxy/service"
	"ai_proxy/service/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

type Response struct {
	Code    string
	Message string
	Data    interface{}
}

func Proxy(c *gin.Context) {
	path := c.Request.URL.Path
	l := logrus.WithField("path", path)
	opt := &common.ReqPayloadOption{}

	if strings.ToUpper(c.Request.Method) == "POST" {
		postBody := make(map[string]interface{})
		if err := c.Bind(&postBody); err != nil {
			l.Error("Bind err : ", err)
			return
		}
		opt.Body = postBody
	}
	logrus.Infof("opt : %+v", opt)
	if err := common.ProxyCall(c, opt); err != nil {
		l.Error("common.GinProxy err : ", err)
		c.JSON(200, Response{
			Code:    "-1",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	if path == config.ConfigInfo.AIDispatcherNodeRegisterPath {
		service.StartHealthCheckReport(opt.Body)
	}
	if path == config.ConfigInfo.AIDispatcherNodeUnRegisterPath {
		service.StopHealthCheckReport()
	}
}
