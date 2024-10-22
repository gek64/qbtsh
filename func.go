package main

import (
	"github.com/gek64/gek/gApp"
	"github.com/gek64/gek/gDownloader"
	"os"
)

func installBinaryFile(tagName string) (err error) {
	downloadURL, err := GetDownloadURL(tagName)
	if err != nil {
		return err
	}
	err = gDownloader.Download(downloadURL, "/usr/local/bin/qbittorrent-nox", "")
	if err != nil {
		return err
	}
	return os.Chmod("/usr/local/bin/qbittorrent-nox", 0755)
}

func installService() (err error) {
	// 获取初始化系统名字,对应的服务内容
	initSystem, serviceContent, err := GetService()
	if err != nil {
		return err
	}
	// 初始化服务
	service, err := gApp.NewService(initSystem, "qbittorrent.service", serviceContent)
	if err != nil {
		return err
	}
	// 服务安装
	err = service.Install()
	if err != nil {
		return err
	}
	// 服务启动,开启自启
	return service.Load()
}

func uninstallBinaryFile() (err error) {
	err = os.RemoveAll("/usr/local/etc/qBittorrent")
	if err != nil {
		return err
	}
	return os.RemoveAll("/usr/local/bin/qbittorrent")
}

func uninstallService() (err error) {
	// 获取初始化系统名字,对应的服务内容
	initSystem, serviceContent, err := GetService()
	if err != nil {
		return err
	}
	// 初始化服务
	service, err := gApp.NewService(initSystem, "qbittorrent.service", serviceContent)
	if err != nil {
		return err
	}
	// 服务停止,关闭自启
	err = service.Unload()
	if err != nil {
		return err
	}
	// 服务卸载
	return service.Uninstall()
}

func updateBinaryFile(tagName string) (err error) {
	// 获取初始化系统名字,对应的服务内容
	initSystem, serviceContent, err := GetService()
	if err != nil {
		return err
	}
	// 初始化服务
	service, err := gApp.NewService(initSystem, "qbittorrent.service", serviceContent)
	if err != nil {
		return err
	}
	// 服务停止,关闭自启
	err = service.Unload()
	if err != nil {
		return err
	}
	// 执行新的二进制文件安装
	err = installBinaryFile(tagName)
	if err != nil {
		return err
	}
	// 服务启动,开启自启
	return service.Load()
}

func updateService() (err error) {
	// 获取初始化系统名字,对应的服务内容
	initSystem, serviceContent, err := GetService()
	if err != nil {
		return err
	}
	// 初始化服务
	service, err := gApp.NewService(initSystem, "qbittorrent.service", serviceContent)
	if err != nil {
		return err
	}
	// 服务停止,关闭自启
	err = service.Unload()
	if err != nil {
		return err
	}
	// 新的服务安装
	err = service.Install()
	if err != nil {
		return err
	}
	// 服务启动,开启自启
	return service.Load()
}

func reloadService() (err error) {
	// 获取初始化系统名字,对应的服务内容
	initSystem, serviceContent, err := GetService()
	if err != nil {
		return err
	}
	// 初始化服务
	service, err := gApp.NewService(initSystem, "qbittorrent.service", serviceContent)
	if err != nil {
		return err
	}
	// 服务重载
	return service.Reload()
}
