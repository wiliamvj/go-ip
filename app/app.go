package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Return app done for running
func Gen() *cli.App {
	app := cli.NewApp()

	app.Name = "Go IP Search"
	app.Usage = "Search IP and servers by command line"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for ips on the web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
