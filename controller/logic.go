package controller

import "github.com/yanshicheng/ikubeops-cli/global"

func (l *LogicService) Run() error {
	// TODO 待完成

	l.L.Info(global.C.Cli.Demo)
	return nil
}
