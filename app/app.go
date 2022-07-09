package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ips := range ips {
		fmt.Println(ips)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, servers := range servers {
		fmt.Println(servers.Host)
	}
}

// Return command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search name and ip address of an server"
	app.Version = "0.0.1"

	flags := []cli.Flag{
		cli.StringFlag{Name: "host", Usage: "Hostname or ip address of an server"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search Ip address of an host",
			Flags: flags,
			Action: func(c *cli.Context) {
				searchIps(c)
			},
		},
		{
			Name:  "servers",
			Usage: "Search server names of an host",
			Flags: flags,
			Action: func(c *cli.Context) {
				searchServers(c)
			},
		},
	}
	return app
}
