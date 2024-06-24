package main

import (
	"ai_proxy/config"
	"ai_proxy/router"
	"ai_proxy/service/common"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var mode string

func init() {
	flag.StringVar(&mode, "mode", "debug", "mode")
}

func main() {
	flag.Parse()
	logrus.Info("mode : ", mode)
	if mode != gin.DebugMode && mode != gin.TestMode && mode != gin.ReleaseMode {
		panic(fmt.Sprintf("invalid mode : %s", mode))
	}

	config.Init(mode)
	common.Init()

	gin.SetMode(mode)
	r := gin.Default()
	router.Init(r)

	r.Run(config.ConfigInfo.ListenAddr)
}
