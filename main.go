package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli"
)

func main() {

	// Setup our Ctrl+C handler
	SetupCloseHandler()

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

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
