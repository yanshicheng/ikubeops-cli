package controller

import (
	"github.com/yanshicheng/ikubeops-cli/global"
	logger "github.com/yanshicheng/ikubeops-cli/settings/ikube-logger"
	"github.com/yanshicheng/ikubeops-cli/version"
)

type LogicService struct {
	service
	L *logger.Logger
}

func NewLogicService() *LogicService {
	version.Info()
	return &LogicService{
		L: global.L.Named("ikube-cli"),
	}
}
