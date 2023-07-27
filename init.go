package main

import (
    "fmt"
    "github.com/gek64/gek/gApp"
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

    // 检查系统中的init system
    _, initSystemBin := gApp.CheckInitSystem()

    // 服务初始化
    switch initSystemBin {
    case gApp.InitSystem["systemd"]:
        bytes, err := container.ReadFile("configs/qbittorrent.service")
        if err != nil {
            return err
        }
        *s = gApp.NewService(cc.Service.Name, string(bytes))
    case gApp.InitSystem["openrc"]:
        return fmt.Errorf("openrc does not yet support")
    case gApp.InitSystem["rc.d"]:
        bytes, err := container.ReadFile("configs/qbittorrent")
        if err != nil {
            return err
        }
        *s = gApp.NewService(cc.Service.Name, string(bytes))
    default:
        var supportInitSystemListString string
        for key := range gApp.InitSystem {
            supportInitSystemListString = supportInitSystemListString + ", " + key
        }
        return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
    }

    return nil
}
