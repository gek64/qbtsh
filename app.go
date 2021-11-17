package main

import (
	"fmt"
	"gek_exec"
	"gek_file"
	"gek_service"
	"gek_toolbox"
	"os"
	"os/exec"
)

var (
	// 工具链
	toolbox = []string{"unzip", "curl", "systemctl", "systemd"}
	// 应用安装目录
	installLocation = "/usr/local/bin/"
	// 应用下载文件存储的临时目录
	tempLocation = "/tmp/qbt_installer/"
	// 应用下载文件中需要的解压的文件
	downloadFileName = []string{"x86_64-qbittorrent-nox", "qbittorrent-nox"}
	// 目标文件名
	targetFileName = "qbittorrent-nox"
	// 默认的原版 qbittorrent-nox
	qbtRepo = "userdocs/qbittorrent-nox-static"
	qbtList = map[string]string{
		"linux_amd64": "x86_64-qbittorrent-nox",
		"linux_arm64": "aarch64-qbittorrent-nox",
		"linux_arm":   "armv7-qbittorrent-nox",
	}
	// qBittorrent Enhanced Edition
	qbteeRepo = "c0re100/qBittorrent-Enhanced-Edition"
	qbteeList = map[string]string{
		"linux_amd64":  "nox_x86_64-linux-musl_static",
		"linux_arm64":  "nox_aarch64-linux-musl_static",
		"linux_arm":    "nox_arm-linux-musleabi_static",
		"linux_mips64": "nox_mips64-linux-musl_static",
		"linux_mips":   "nox_mips-linux-musl_static",
		"linux_mipsle": "nox_mipsel-linux-musl_static",
	}
	// 服务名称
	serviceName = "qbittorrent-nox.service"
	// 服务内容
	serviceContent = `[Unit]
Description=qBittorrent Daemon Service
Wants=network-online.target
After=network-online.target nss-lookup.target

[Service]
Type=exec
User=root
ExecStart=/usr/local/bin/qbittorrent-nox
Restart=on-failure
SyslogIdentifier=qbittorrent-nox

[Install]
WantedBy=multi-user.target`
)

// 主要功能函数
// 安装
func install(repo string, repoList map[string]string, needExtract bool) (err error) {
	// 需要文件未安装
	exist, _, _ := gek_file.Exist(installLocation + targetFileName)
	if exist {
		return fmt.Errorf("app is already installed")
	}
	// 下载应用
	err = downloadApp(repo, repoList)
	if err != nil {
		return err
	}
	// 处理应用
	err = processApp(needExtract)
	if err != nil {
		return err
	}
	// 安装服务
	err = installService()
	if err != nil {
		return err
	}
	return nil
}

// 卸载
func uninstall() (err error) {
	// 重置服务
	err = resetApp()
	if err != nil {
		return
	}
	// 卸载服务
	err = uninstallService()
	if err != nil {
		return err
	}
	return nil
}

// 更新
func update(repo string, repoList map[string]string, needExtract bool) (err error) {
	// 需要应用已安装
	exist, _, _ := gek_file.Exist(installLocation + targetFileName)
	if !exist {
		return fmt.Errorf("app is not installed")
	}
	// 下载应用
	err = downloadApp(repo, repoList)
	if err != nil {
		return err
	}
	// 处理应用
	err = processApp(needExtract)
	if err != nil {
		return err
	}
	// 重载应用
	err = reload()
	if err != nil {
		return err
	}
	return nil
}

// 重载
func reload() (err error) {
	// 重载服务
	err = reloadService()
	if err != nil {
		return err
	}
	return nil
}

// 功能实现函数
// 下载应用
func downloadApp(repo string, repoList map[string]string) (err error) {
	// 工作路径
	originDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// 建立临时文件夹
	err = gek_file.CreateDir(tempLocation)
	if err != nil {
		return err
	}
	// 获取下载链接
	downloadLink, err := gek_toolbox.GetDownloadLink(repo, repoList)
	if err != nil {
		return err
	}
	// cd 临时文件夹
	err = os.Chdir(tempLocation)
	if err != nil {
		return err
	}
	// 下载文件到临时文件夹
	err = gek_exec.Run("curl -LOJ " + downloadLink)
	if err != nil {
		return err
	}
	// cd 工作路径
	err = os.Chdir(originDir)
	if err != nil {
		return err
	}
	return nil
}

// 处理应用下载文件
func processApp(needExtract bool) (err error) {
	// 如果不存在安装文件夹,则创建
	exist, _, err := gek_file.Exist(installLocation)
	if err != nil || !exist {
		err := gek_file.CreateDir(installLocation)
		if err != nil {
			return err
		}
	}
	if needExtract {
		err = gek_exec.Run(exec.Command("unzip", "-o", tempLocation+"*.zip", downloadFileName[1], "-d", installLocation))
		if err != nil {
			return err
		}
	} else {
		err = gek_exec.Run(exec.Command("mv", tempLocation+downloadFileName[0], installLocation+targetFileName))
		if err != nil {
			return err
		}
	}
	// 赋权755
	err = os.Chmod(installLocation+targetFileName, 755)
	if err != nil {
		return err
	}
	// 删除临时文件夹
	err = os.RemoveAll(tempLocation)
	if err != nil {
		return err
	}
	return nil
}

// 清理应用文件
func resetApp() (err error) {
	// 删除应用文件
	err = os.RemoveAll(installLocation + targetFileName)
	if err != nil {
		return err
	}
	// 删除配置文件及缓存
	err = gek_exec.Run(exec.Command("rm", "-rf", "/root/.cache/qBittorrent/", "/root/.config/qBittorrent/", "/root/.local/share/qBittorrent/"))
	if err != nil {
		return err
	}
	return nil
}

// 安装服务
func installService() (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 安装服务
	err = service.Install()
	if err != nil {
		return err
	}
	// 查看服务状态
	err = service.Status()
	if err != nil {
		return err
	}
	return nil
}

// 卸载服务
func uninstallService() (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 卸载服务
	err = service.Uninstall()
	if err != nil {
		return err
	}
	return nil
}

// 重载服务
func reloadService() (err error) {
	service := gek_service.NewService(serviceName, serviceContent)
	// 重载服务
	err = service.Reload()
	if err != nil {
		return err
	}
	// 查看服务状态
	err = service.Status()
	if err != nil {
		return err
	}
	return nil
}
