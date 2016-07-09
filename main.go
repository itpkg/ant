package main

import (
	"os"

	"github.com/itpkg/ant/claw"
	"github.com/itpkg/ant/minify"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ant"
	app.Usage = "claw a site to epub book."
	app.Version = "v20160708"
	app.Commands = []cli.Command{
		claw.Command,
		minify.Command,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
