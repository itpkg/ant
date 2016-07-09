package claw

import "github.com/urfave/cli"

//Command claw
var Command = cli.Command{

	Name:    "claw",
	Aliases: []string{"c"},
	Usage:   "get whole wesite pages",
	Flags: []cli.Flag{
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
	},
	Action: func(c *cli.Context) error {
		url := c.String("url")
		if len(url) == 0 {
			cli.ShowCommandHelp(c, "claw")
			return nil
		}
		return claw(url)
	},
}
