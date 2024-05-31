package main

import (
	"embed"
	"errors"
	"net/url"
	"os/exec"
	"runtime"
)

//go:embed configs/*
var container embed.FS

func GetDownloadURL() (downloadURL string, err error) {
	baseURL := "https://github.com/userdocs/qbittorrent-nox-static/releases/latest/download/"
	assetName := ""
	switch runtime.GOARCH {
	case "amd64":
		assetName = "x86_64-qbittorrent-nox"
	case "386":
		assetName = "x86-qbittorrent-nox"
	case "arm":
		assetName = "armv7-qbittorrent-nox"
	case "arm64":
		assetName = "aarch64-qbittorrent-nox"
	}
	return url.JoinPath(baseURL, assetName)
}

func GetService() (initSystem string, serviceContent []byte, err error) {
	serviceFile := ""
	// systemd
	_, err = exec.LookPath("systemctl")
	if err == nil {
		serviceFile = "configs/qbittorrent.service"
		initSystem = "systemd"
	}
	// alpine openrc
	_, err = exec.LookPath("openrc")
	if err == nil {
		serviceFile = "configs/qbittorrent.openrc"
		initSystem = "openrc"
	}
	// freebsd rc.d
	_, err = exec.LookPath("rcorder")
	if err == nil {
		serviceFile = "configs/qbittorrent.rcd"
		initSystem = "rc.d"
	}

	// 找不到初始化系统返回错误
	if initSystem == "" {
		return "", nil, errors.New("init system not found")
	}

	// 读取文件并返回文件内容
	bytes, err := container.ReadFile(serviceFile)
	if err != nil {
		return initSystem, nil, err
	}
	return initSystem, bytes, nil
}
