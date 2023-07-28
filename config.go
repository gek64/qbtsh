package main

import (
    "embed"
    "encoding/json"
    "fmt"
    "github.com/gek64/gek/gApp"
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
//
//go:embed configs/*
var container embed.FS

func initConf() (err error) {
    var configBytes []byte
    var configPath string

    // 检查系统中的init system
    _, initSystemBin := gApp.CheckInitSystem()

    switch runtime.GOOS {
    case gApp.SupportedOS[0]:
        switch initSystemBin {
        case gApp.InitSystem["systemd"]:
            configPath = "configs/linux_systemd.json"
        case gApp.InitSystem["openrc"]:
            configPath = "configs/linux_openrc.json"
        default:
            var supportInitSystemListString string
            for key := range gApp.InitSystem {
                supportInitSystemListString = supportInitSystemListString + ", " + key
            }
            return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
        }
    case gApp.SupportedOS[1]:
        switch initSystemBin {
        case gApp.InitSystem["rc.d"]:
            configPath = "configs/freebsd_rcd.json"
            return fmt.Errorf("freebsd does not yet support")
        default:
            var supportInitSystemListString string
            for key := range gApp.InitSystem {
                supportInitSystemListString = supportInitSystemListString + ", " + key
            }
            return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
        }
    }

    configBytes, err = container.ReadFile(configPath)
    if err != nil {
        return err
    }

    err = json.Unmarshal(configBytes, &cc)
    if err != nil {
        return err
    }

    return nil
}
