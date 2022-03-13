package main

import (
	"embed"
	"encoding/json"
	"gek_app"
	"runtime"
)

type CC struct {
	Toolbox     []string    `json:"toolbox"`
	Application Application `json:"application"`
	Config      Config      `json:"config"`
	Service     Service     `json:"service"`
}

type Application struct {
	Repo                    string            `json:"repo"`
	RepoList                map[string]string `json:"repoList"`
	File                    []string          `json:"file"`
	Location                string            `json:"location"`
	UninstallDeleteLocation bool              `json:"uninstallDeleteLocation"`
}
type Config struct {
	Location                string `json:"location"`
	UninstallDeleteLocation bool   `json:"uninstallDeleteLocation"`
}
type Service struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// 存储配置文件及服务文件
//go:embed config/* service/*
var container embed.FS

func initConf() (err error) {
	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		bytes, err := container.ReadFile("config/linux.json")
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &cc)
		if err != nil {
			return err
		}
	}
	return nil
}
