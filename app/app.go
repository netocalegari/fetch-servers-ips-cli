package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns the cli app ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI app"
	app.Usage = "Search for IPs and server name on the web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs on the web",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search for server name on the web",
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
		log.Println(ip)
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
