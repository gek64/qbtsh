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
	cliHelp      bool
	cliVersion   bool
)

func init() {
	flag.BoolVar(&cliInstall, "install", false, "install")
	flag.BoolVar(&cliUninstall, "uninstall", false, "uninstall")
	flag.BoolVar(&cliUpdate, "update", false, "update")
	flag.BoolVar(&cliReload, "reload", false, "reload")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    qbtsh [Commands]

Command:
    -install    : Install
    -uninstall  : Uninstall
    -update     : Update
    -reload     : Reload
    -h          : Show help
    -v          : Show version

Example:
    1) qbtsh -install     : Install qBittorrent-nox
    3) qbtsh -update      : Update qBittorrent-nox
    5) qbtsh -uninstall   : Remove config,cache and uninstall qBittorrent-nox
    6) qbtsh -reload      : Reload service`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.03")
		os.Exit(0)
	}

	// 检查运行库是否完整
	err := gek_toolbox.CheckToolbox(toolbox)
	if err != nil {
		log.Fatal(err)
	}

	err = initNetwork()
	if err != nil {
		fmt.Println("init fail")
		log.Fatal(err)
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Now add aria2c,wget and build-in downloader support
  1.02:
    - Rewrite download function
  1.03:
    - Rewrite all code
    - Fix the description in the error message`
	fmt.Println(versionInfo)
}

func main() {
	if cliInstall {
		err := install()
		if err != nil {
			log.Panicln(err)
		}
	}
	if cliUpdate {
		err := update()
		if err != nil {
			log.Panicln(err)
		}
	}
	if cliUninstall {
		err := uninstall()
		if err != nil {
			log.Panicln(err)
		}
	}
	if cliReload {
		err := reload()
		if err != nil {
			log.Panicln(err)
		}
	}
}
