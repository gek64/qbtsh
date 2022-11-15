package main

import (
	"github.com/gek64/gek/gApp"
	"runtime"
)

var (
	app         gApp.Application
	config      gApp.Config
	service     gApp.Service
	tempFolder  = "/tmp/qbt_installer"
	needExtract = false
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

	*a, err = gApp.NewApplicationFromGithub(cc.Application.File, cc.Application.Repo, cc.Application.RepoList, cc.Application.Location, cc.Application.UninstallDeleteLocation, tempFolder)
	if err != nil {
		return err
	}

	// 配置初始化
	*c = gApp.NewConfig("", "", cc.Config.Location, cc.Config.UninstallDeleteLocation)

	// 服务初始化
	switch runtime.GOOS {
	case gApp.SupportedOS[0]:
		bytes, err := container.ReadFile("configs/qbittorrent.service")
		if err != nil {
			return err
		}
		*s = gApp.NewService(cc.Service.Name, string(bytes))
	case gApp.SupportedOS[1]:
		bytes, err := container.ReadFile("configs/qbittorrent")
		if err != nil {
			return err
		}
		*s = gApp.NewService(cc.Service.Name, string(bytes))
	}

	return nil
}
