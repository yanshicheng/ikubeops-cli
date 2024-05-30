package global

import (
	logger "github.com/yanshicheng/ikubeops-cli/settings/ikube-logger"
)

const (
	EnvPrefix                 = "IKUBEOPS_"
	ConfigModeFile ConfigMode = "file"
	ConfigModeEnv  ConfigMode = "env"
)

type ConfigMode string

var (
	// Config 全局配置
	C    *Config = NewDefaultConfig()
	L    *logger.Logger
	LSys *logger.Logger
)
