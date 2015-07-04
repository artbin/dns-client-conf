package main

import (
	"log"
	"net"
	"os"

	"github.com/ArtemKulyabin/dns-client-conf"
	"github.com/ArtemKulyabin/dns-client-conf/debugmode"
	"github.com/ArtemKulyabin/dns-client-conf/helpers"

	"github.com/codegangsta/cli"
)

var dnsconf = dnsclientconf.NewDNSConfigurator()

func getNameServers(c *cli.Context) {
	addrs, err := dnsconf.GetNameServers()
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		println(addr.String())
	}
}

func addNameServers(c *cli.Context) {
	var addrs []net.IP
	for _, addr := range c.Args() {
		err := helpers.CheckIP(addr)
		if err != nil {
			log.Fatal(err)
		}
		addrs = append(addrs, net.ParseIP(addr))
	}
	err := dnsconf.AddNameServers(addrs)
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
		if c.IsSet("interface") {
			iface, err := net.InterfaceByName(c.String("interface"))
			if err != nil {
				log.Fatal(err)
			}
			dnsconf.SetInterface(iface)
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "activate debug mode",
		},
		cli.StringFlag{
			Name:  "interface",
			Usage: "network interface name",
		},
	}

	app.Run(os.Args)
}
