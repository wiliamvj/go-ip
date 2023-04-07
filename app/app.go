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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for ips on the web",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search for servers on the web",
			Flags:  flags,
			Action: searchServers,
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

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
