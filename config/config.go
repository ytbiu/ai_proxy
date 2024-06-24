package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	DevMode     = "dev"
	TestMode    = "test"
	ReleaseMode = "release"
)

var ConfigInfo Info

type Info struct {
	ListenAddr                  string
	AIDispatcherNodeProxyAddr   string
	AITaskExecutorNodeProxyAddr string

	AIDispatcherNodeProxyAPIPath   []string
	AITaskExecutorNodeProxyAPIPath []string
}

func Init(mode string) {
	cfgFile := "config-dev.yaml"
	if mode == gin.TestMode {
		cfgFile = "config-test.yaml"
	}
	if mode == gin.ReleaseMode {
		cfgFile = "config-release.yaml"
	}
	viper.SetConfigName(cfgFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	MustNilErr(err, "ReadInConfig failed")

	err = viper.Unmarshal(&ConfigInfo)
	MustNilErr(err, "config Unmarshal failed")
}

func MustNilErr(err error, msg string) {
	if err != nil {
		logrus.Fatalf("msg : %s, err : %s", msg, err)
	}
}
