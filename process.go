package main

import (
	"log"
)

// 安装
func install() (err error) {
	// 应用安装
	err = app.Install(tempFolder, true, needExtract)
	if err != nil {
		return err
	}
	// 服务安装
	err = service.Install()
	if err != nil {
		return err
	}

	return nil
}

// 卸载
func uninstall() (err error) {
	// 停止应用+卸载服务
	err = service.Uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载配置文件
	err = config.Uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载应用
	err = app.Uninstall()
	if err != nil {
		log.Println(err)
	}

	return nil
}

// 更新
func update() (err error) {
	// 应用卸载
	err = app.Uninstall()
	if err != nil {
		return err
	}
	// 应用安装
	err = app.Install(tempFolder, true, needExtract)
	if err != nil {
		return err
	}

	// 服务重启
	err = service.Restart()
	if err != nil {
		return err
	}

	return nil
}

// 重载
func reload() (err error) {
	// 服务重启
	err = service.Restart()
	if err != nil {
		return err
	}

	return nil
}
