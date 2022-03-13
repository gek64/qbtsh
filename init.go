package main

import (
	"gek_app"
	"runtime"
)

var (
	app         gek_app.Application
	config      gek_app.Config
	service     gek_app.Service
	tempFolder  string = "/tmp/qbt_installer"
	needExtract bool   = false
	cc          CC
)

func initApp() (err error) {
	var a = &app
	var c = &config
	var s = &service

	err = initConf()
	if err != nil {
		return err
	}

	// 应用初始化

	*a, err = gek_app.NewApplicationFromGithub(cc.Application.File, cc.Application.Repo, cc.Application.RepoList, cc.Application.Location, cc.Application.UninstallDeleteLocation, tempFolder)
	if err != nil {
		return err
	}

	// 配置初始化
	*c = gek_app.NewConfig("", "", cc.Config.Location, cc.Config.UninstallDeleteLocation)

	// 服务初始化
	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		bytes, err := container.ReadFile("service/qbittorrent.service")
		if err != nil {
			return err
		}
		*s = gek_app.NewService(cc.Service.Name, string(bytes))
	case gek_app.SupportedOS[1]:
		bytes, err := container.ReadFile("service/qbittorrent")
		if err != nil {
			return err
		}
		*s = gek_app.NewService(cc.Service.Name, string(bytes))
	}

	return nil
}
