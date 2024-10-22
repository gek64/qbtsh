package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var tagName string

	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "Install qBittorrent",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "tag",
					Aliases:     []string{"t"},
					Usage:       "set qBittorrent tag name",
					Destination: &tagName,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				err = installBinaryFile(tagName)
				if err != nil {
					return err
				}
				return installService()
			},
		},
		{
			Name:  "uninstall",
			Usage: "Remove config,cache and uninstall qBittorrent",
			Action: func(ctx *cli.Context) (err error) {
				err = uninstallService()
				if err != nil {
					return err
				}
				return uninstallBinaryFile()
			},
		},
		{
			Name:  "update",
			Usage: "Update qBittorrent",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "tag",
					Aliases:     []string{"t"},
					Usage:       "set qBittorrent tag name",
					Destination: &tagName,
				},
			},
			Action: func(ctx *cli.Context) (err error) {
				err = updateBinaryFile(tagName)
				if err != nil {
					return err
				}
				return updateService()
			},
		},
		{
			Name:  "reload",
			Usage: "Reload service",
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
		Usage:    "qBittorrent quick install tool",
		Version:  "v2.01",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
