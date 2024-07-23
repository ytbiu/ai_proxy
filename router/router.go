package router

import (
	"ai_proxy/config"
	"ai_proxy/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	for _, path := range config.ConfigInfo.AITaskExecutorNodeProxyAPIPath {
		r.Any(path, controller.Proxy)
	}
}
