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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "find address",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},

			Action: func(c *cli.Context) {
				host := c.String("host")
				ips, err := net.LookupIP(host)

				if err != nil {
					log.Fatal(err)
				}

				for _, ip := range ips {
					fmt.Println(ip)
				}
			},
		},
	}
	return app
}
