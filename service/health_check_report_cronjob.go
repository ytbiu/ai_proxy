package service

import (
	"ai_proxy/config"
	"ai_proxy/service/common"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
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
		for {
			if common.FileExists(config.ConfigInfo.RegisterDataFile) {
				if reportInfo == nil {
					common.LoadFile(&reportInfo)
				}

				err := common.Post(config.ConfigInfo.HealthCheckServiceReportUrl, nil, map[string]interface{}{
					"node_id": reportInfo.NodeId,
					"project": reportInfo.Project,
					"models":  reportInfo.Models,
				})
				if err != nil {
					logrus.WithField("node_id", reportInfo.NodeId).
						WithField("project", reportInfo.Project).
						WithField("models", reportInfo.Models).
						Errorf("health check report err : %s", err)
				}
				logrus.Info("health check reported")
				time.Sleep(periodSecond)
			}
		}
	}()
}
