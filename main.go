package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Go Wait toolkit"
	app.Usage = "wait hostname[:port][ hostname:[port] ...]"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "t,timeout",
			Usage: "Interval after waiting for each port is considered as timeout(seconds, default: 60)",
			Value: 60,
		},
	}
	app.Action = func(c *cli.Context) error {
		args := c.Args()

		timeout := time.Duration(c.Int("timeout")) * time.Second

		if err := waitForServices(args, timeout); err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	app.Run(os.Args)
}
