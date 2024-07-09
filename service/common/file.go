package common

import (
	"ai_proxy/config"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func WriteDataToFile(data interface{}, filePath string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func LoadFile(result interface{}) {
	file, err := os.Open(config.ConfigInfo.RegisterDataFile)
	if err != nil {
		logrus.Error("open file err :", err)
		return
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		logrus.Error("read file err :", err)
		return
	}

	if err := json.Unmarshal(contents, result); err != nil {
		logrus.Error("json unmarshal:", err)
	}
}
