package service

import (
	"ai_proxy/config"
	"ai_proxy/service/common"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

type ReportInfo struct {
	NodeId  string
	Project string
	Models  []RegisterModel
}

var reportInfo *ReportInfo

type RegisterPayload struct {
	Project string
	Models  []RegisterModel
}

type RegisterModel struct {
	Model string
}

var nodeId string

func InitNodeId() {
	var response struct {
		PeerId string `json:"peer_id"`
	}
	err := common.Get(config.ConfigInfo.GetPeerIdUrl(), &response)
	if err != nil {
		logrus.Errorf("get perr id err : %s", err)
		return
	}
	nodeId = response.PeerId
	logrus.Info("get node Id : ", nodeId)
}

func StartHealthCheckReport(payloadMap map[string]interface{}) {
	var payload RegisterPayload
	if err := mapstructure.Decode(payloadMap, &payload); err != nil {
		logrus.Errorf("copier err : %s", err)
		return
	}

	reportInfo = &ReportInfo{
		NodeId:  nodeId,
		Project: payload.Project,
		Models:  payload.Models,
	}

	err := common.WriteDataToFile(reportInfo, config.ConfigInfo.RegisterDataFile)
	if err != nil {
		logrus.Errorf("common.WriteDataToFile err : %s. RegisterDataFile : %s", err, config.ConfigInfo.RegisterDataFile)
		return
	}
}

func StopHealthCheckReport() {
	err := common.DeleteFile(config.ConfigInfo.RegisterDataFile)
	if err != nil {
		logrus.Errorf("common.DeleteFile err : %s. RegisterDataFile : %s", err, config.ConfigInfo.RegisterDataFile)
		return
	}
	reportInfo = &ReportInfo{}
}

func HealthCheckReportCronJob() {
	go func() {
		periodSecond := time.Duration(config.ConfigInfo.HealthCheckReportPeriodSeconds) * time.Second
		logrus.Infof("health check report will exec every %s after register", periodSecond)
		reportCronJob(periodSecond)
	}()
}

func reportCronJob(periodSecond time.Duration) {
	// wait for ws server start
	time.Sleep(time.Second * 5)

	u := url.URL{Scheme: "ws", Host: config.ConfigInfo.HealthCheckServiceReportAddr, Path: config.ConfigInfo.HealthCheckServiceReportPath}
	logrus.Infof("Connecting to %s", u.String())

	// dial WebSocket server
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		if common.FileExists(config.ConfigInfo.RegisterDataFile) {
			if reportInfo == nil {
				common.LoadFile(&reportInfo)
			}
		}

		jsonData, _ := json.Marshal(reportInfo)
		err := c.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			logrus.WithField("node_id", reportInfo.NodeId).
				WithField("project", reportInfo.Project).
				WithField("models", reportInfo.Models).
				Errorf("health check report err : %s", err)
			TryToReconnect(c)
		}
		time.Sleep(periodSecond)
	}
}

func TryToReconnect(c *websocket.Conn) {
	if c.WriteMessage(websocket.PingMessage, []byte{}) != nil {
		u := url.URL{Scheme: "ws", Host: config.ConfigInfo.HealthCheckServiceReportAddr, Path: config.ConfigInfo.HealthCheckServiceReportPath}
		logrus.Infof("Connecting to %s", u.String())

		// dial WebSocket server
		reConnect, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			logrus.Fatal("dial:", err)
		}
		c = reConnect
		return
	}
}
