package common

import "ai_proxy/config"

var Path2ProxyAddr = map[string]string{}

func Init() {
	// docs:https://hieosqe4lol.feishu.cn/docx/Rk8pd1w9CoottHxOn8ZcwifwnBc
	for _, path := range config.ConfigInfo.AIDispatcherNodeProxyAPIPath {
		Path2ProxyAddr[path] = config.ConfigInfo.AIDispatcherNodeProxyAddr
	}
	// docs:https://hieosqe4lol.feishu.cn/docx/EqPedXT8Io7EYMxtdgucgfqRngb
	for _, path := range config.ConfigInfo.AITaskExecutorNodeProxyAPIPath {
		Path2ProxyAddr[path] = config.ConfigInfo.AITaskExecutorNodeProxyAddr
	}
}

func ShouldDirectProxy(path string) bool {
	_, found := Path2ProxyAddr[path]
	return found
}
