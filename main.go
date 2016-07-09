package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ant"
	app.Usage = "claw a site to epub book."
	app.Version = "v20160708"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Value: "",
			Usage: "main page's url",
		},
		cli.UintFlag{
			Name:  "timeout, t",
			Value: 5,
			Usage: "timeout by second",
		},
		cli.Uint64Flag{
			Name:  "retry, r",
			Value: 3,
			Usage: "retry times",
		},
		cli.Float64Flag{
			Name:  "sleep, s",
			Value: 0.2,
			Usage: "sleep by second",
		},
		cli.StringFlag{
			Name:  "book, b",
			Value: "",
			Usage: "book's name, default by pate's title",
		},
		cli.StringFlag{
			Name:  "format, f",
			Value: "epub",
			Usage: "book's format, like: epub",
		},
	}
	app.Action = func(c *cli.Context) error {
		url := c.String("url")
		if len(url) == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		return claw(url)
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
