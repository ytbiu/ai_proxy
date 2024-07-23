package common

import "ai_proxy/config"

var Path2ProxyAddr = map[string]string{}

func Init() {
	// docs:https://hieosqe4lol.feishu.cn/docx/EqPedXT8Io7EYMxtdgucgfqRngb
	for _, path := range config.ConfigInfo.AITaskExecutorNodeProxyAPIPath {
		Path2ProxyAddr[path] = config.ConfigInfo.AITaskExecutorNodeProxyAddr
	}
}
