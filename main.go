package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install qBittorrent-nox",
			Action: func(ctx *cli.Context) (err error) {
				err = installBinaryFile()
				if err != nil {
					return err
				}
				return installService()
			},
		},
		{
			Name:    "uninstall",
			Aliases: []string{"u"},
			Usage:   "Remove config,cache and uninstall qBittorrent-nox",
			Action: func(ctx *cli.Context) (err error) {
				err = uninstallService()
				if err != nil {
					return err
				}
				return uninstallBinaryFile()
			},
		},
		{
			Name:    "update",
			Aliases: []string{"up"},
			Usage:   "Update qBittorrent-nox",
			Action: func(ctx *cli.Context) (err error) {
				err = updateBinaryFile()
				if err != nil {
					return err
				}
				return updateService()
			},
		},
		{
			Name:    "reload",
			Aliases: []string{"r"},
			Usage:   "Reload service",
			Action: func(ctx *cli.Context) (err error) {
				return reloadService()
			},
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:    "qBittorrent-nox quick install tool",
		Version:  "v2.00",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
