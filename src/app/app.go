package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// return command line app
func Start() *cli.App {
	app := cli.NewApp()
	app.Name = "app"
	app.Usage = "find web servers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "find address",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servers",
			Usage:  "find servers",
			Flags:  flags,
			Action: findServers,
		},
	}
	return app
}

func findServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func findIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
