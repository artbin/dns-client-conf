package main

import (
	"log"
	"os"
	"strings"

	"github.com/ArtemKulyabin/dns-client-conf"

	"github.com/codegangsta/cli"
)

func main() {

	dnsconf := dnsclientconf.NewDNSConfigurator()

	app := cli.NewApp()
	app.Name = "dns-client-conf"
	app.Usage = "portable configuration tool for work with operating system dns client settings"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "show",
			Usage: "show the current name servers",
			Action: func(c *cli.Context) {
				addrs, err := dnsconf.GetNameServers()
				if err != nil {
					log.Fatal(err)
				}
				println(strings.Join(addrs, "\n"))
			},
		},
		{
			Name:  "add",
			Usage: "add name servers",
			Action: func(c *cli.Context) {
				err := dnsconf.AddNameServers(c.Args())
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("DNS servers added successfully")
				}
			},
		},
		{
			Name:  "dhcp",
			Usage: "receive dns addresses by DHCP",
			Action: func(c *cli.Context) {
				err := dnsconf.DHCPNameServers()
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("DNS servers received by DCHP successfully")
				}
			},
		},
		{
			Name:  "reload",
			Usage: "reload name servers settings",
			Action: func(c *cli.Context) {
				err := dnsconf.ReloadNameServers()
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("DNS servers reloaded successfully")
				}
			},
		},
	}

	app.Run(os.Args)
}
