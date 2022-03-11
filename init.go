package main

import (
	"fmt"
	"gek_app"
	"runtime"
)

var (
	app         gek_app.Application
	config      gek_app.Config
	service     gek_app.Service
	tempFolder  string = "/tmp/qbt_installer"
	needExtract bool   = false
)

func initNetwork() (err error) {
	var a = &app
	var c = &config
	var s = &service

	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		// 应用初始化
		*a, err = gek_app.NewApplicationFromGithub(bins, repo, repoList, binsLocation, binUninstallDeleteLocationFolder, tempFolder)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = gek_app.NewConfig(configName, configContent, configLocation, configUninstallDeleteLocationFolder)

		// 服务初始化
		*s = gek_app.NewService(serviceName, serviceContent)
	default:
		return fmt.Errorf("unsupported os %s", runtime.GOOS)
	}

	return nil
}
