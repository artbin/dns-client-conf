package main

import (
	"log"
	"os"
	"strings"

	"github.com/ArtemKulyabin/dns-client-conf"
	"github.com/ArtemKulyabin/dns-client-conf/debugmode"

	"github.com/codegangsta/cli"
)

var dnsconf = dnsclientconf.NewDNSConfigurator()

func getNameServers(c *cli.Context) {
	addrs, err := dnsconf.GetNameServers()
	if err != nil {
		log.Fatal(err)
	}
	println(strings.Join(addrs, "\n"))
}

func addNameServers(c *cli.Context) {
	err := dnsconf.AddNameServers(c.Args())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DNS servers added successfully")
	}
}

func dHCPNameServers(c *cli.Context) {
	err := dnsconf.DHCPNameServers()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DNS servers received by DCHP successfully")
	}
}

func reloadNameServers(c *cli.Context) {
	err := dnsconf.ReloadNameServers()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DNS servers reloaded successfully")
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "dns-client-conf"
	app.Usage = "portable configuration tool for work with operating system dns client settings"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:   "show",
			Usage:  "show the current name servers",
			Action: getNameServers,
		},
		{
			Name:   "add",
			Usage:  "add name servers",
			Action: addNameServers,
		},
		{
			Name:   "dhcp",
			Usage:  "receive dns addresses by DHCP",
			Action: dHCPNameServers,
		},
		{
			Name:   "reload",
			Usage:  "reload name servers settings",
			Action: reloadNameServers,
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.IsSet("debug") {
			debugmode.ActivateDebugMode()
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "activate debug mode",
		},
	}

	app.Run(os.Args)
}
