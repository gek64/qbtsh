package main

import (
	"flag"
	"fmt"
	"gek_toolbox"
	"log"
	"os"
)

var (
	cliInstall   bool
	cliUninstall bool
	cliUpdate    bool
	cliReload    bool
	cliEE        bool
	cliHelp      bool
	cliVersion   bool
)

func init() {
	flag.BoolVar(&cliInstall, "I", false, "install")
	flag.BoolVar(&cliUninstall, "R", false, "uninstall")
	flag.BoolVar(&cliUpdate, "U", false, "update")
	flag.BoolVar(&cliReload, "r", false, "reload")
	flag.BoolVar(&cliEE, "ee", false, "choose qBittorrent-nox enhanced edition")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    qbtsh [Options] [Commands]

Options:
    -ee   : Choose qBittorrent-nox enhanced edition

Command:
    -I    : Install
    -R    : Uninstall
    -U    : Update
    -r    : Reload
    -h    : Show help
    -v    : Show version

Example:
    1) qbtsh -I      : Install default qBittorrent-nox
    2) qbtsh -ee -I  : Install qBittorrent-nox enhanced edition
    3) qbtsh -U      : Update default qBittorrent-nox
    4) qbtsh -ee -U  : Update qBittorrent-nox enhanced edition
    5) qbtsh -R      : Remove cache and uninstall qBittorrent-nox
    6) qbtsh -r      : Reload service`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.01")
		os.Exit(0)
	}

	// 检查运行库是否完整
	err := gek_toolbox.CheckToolbox(toolbox)
	if err != nil {
		log.Fatal(err)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Now add aria2c,wget and build-in downloader support`
	fmt.Println(versionInfo)
}

func main() {
	if cliInstall {
		if cliEE {
			err := install(qbteeRepo, qbteeList, true)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := install(qbtRepo, qbtList, false)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if cliUpdate {
		if cliEE {
			err := update(qbteeRepo, qbteeList, true)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := update(qbtRepo, qbtList, false)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if cliUninstall {
		err := uninstall()
		if err != nil {
			log.Println(err)
		}
	}
	if cliReload {
		err := reload()
		if err != nil {
			log.Println(err)
		}
	}
}
