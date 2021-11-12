package main

import (
	"github.com/urfave/cli/v2"
	"launcher/internal"
	"launcher/internal/biz/danmaku"
	"log"
	"os"
)

func main () {
	app := &cli.App{
		Name: "sikong.ren tools services",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "config",
				Aliases: []string { "c" },
				Value: "./configs/server.yml",
				Usage: "server config file",
			},
			&cli.StringFlag{
				Name: "listen",
				Aliases: []string { "l" },
				Value: "",
				Usage: "listen address",
			},
		},
		Action: func(context *cli.Context) error {
			return internal.RunServer(context.String("config"), context.String("listen"))
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name: "danmaku",
				Action: func(context *cli.Context) error {
					return danmaku.RunServer(context.String("config"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "config",
						Aliases: []string { "c" },
						Value: "./configs/server.yml",
						Usage: "server config file",
					},
				},
			},
			&cli.Command{
				Name: "httpserver",
				Aliases: []string{ "server", "http" },
				Action: func(context *cli.Context) error {
					return nil
				},
			},
			&cli.Command{
				Name: "websocket",
				Aliases: []string{ "ws" },
				Action: func(context *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
