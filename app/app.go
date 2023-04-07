package app

import "github.com/urfave/cli"

// Return app done for running
func Gen() *cli.App {
	app := cli.NewApp()

	app.Name = "Go IP Search"
	app.Usage = "Search IP and servers by command line"

	return app
}
