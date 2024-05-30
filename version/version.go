package version

import (
	"fmt"
	"github.com/yanshicheng/ikubeops-cli/global"
)

const (
	IkubeopsDescription = "这是一个 gin 模板项目"
)

var (
	IkubeopsTag         string
	IkubeopsCommit      string
	IkubeopsBranch      string
	IkubeopsBuildTime   string
	IkubeopsGoVersion   string
	IkubeopsUrl         string
	IkubeopsProjectName = "ikubeops-cli"
	IkubeopsConfigType  string
	IkubeopsConfigFile  string
	IkubeopsConfigEtcd  string
)

// FullTagVersion FullVersion show the version info
func FullTagVersion() string {
	version := fmt.Sprintf("Version   : %s\nBuild Time: %s\nGit URL: %s\nGit Branch: %s\nGo Version: %s\n", IkubeopsCommit, IkubeopsBuildTime, IkubeopsUrl, IkubeopsBranch, IkubeopsGoVersion)
	return version
}

// ShortTagVersion Short 版本缩写
func ShortTagVersion() string {
	return fmt.Sprintf("%s[%s %s]", IkubeopsTag, IkubeopsBuildTime, IkubeopsCommit)
}

func GetWebUrl() string {
	addr := fmt.Sprintf("%s:%d", global.C.App.HttpAddr, global.C.App.HttpPort)
	if global.C.App.Tls {
		return fmt.Sprintf("https://%s", addr)
	} else {
		return fmt.Sprintf("http://%s", addr)
	}
}

func GetConfig() string {
	if IkubeopsConfigType == "file" {
		return fmt.Sprintf("File: %s", IkubeopsConfigFile)
	} else if IkubeopsConfigType == "env" {
		return "ENV"
	} else {
		return fmt.Sprintf("ETCD: %s", IkubeopsConfigEtcd)
	}
}

func Info() {
	fmt.Printf(`
欢迎使用: %s
当前版本: %s
配置文件: %s
演示地址: www.ikubeops.com
代码地址: %s
运行地址: %s

`, IkubeopsProjectName, ShortTagVersion(), GetConfig(), IkubeopsUrl, GetWebUrl())
}
