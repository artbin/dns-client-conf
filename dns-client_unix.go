// +build linux freebsd openbsd netbsd

package dnsclientconf

import (
	"github.com/ArtemKulyabin/dns-client-conf/dhclient"
	"github.com/ArtemKulyabin/dns-client-conf/resolvconf"
)

type dNSConfig struct {
	resolvConfigPath, dhclientConfigPath, dhclientConfigPathBackup string
}

func NewDNSConfigurator() DNSConfigurator {
	return &dNSConfig{ResolvConfigPath, DhclientConfigPath, DhclientConfigPathBackup}
}

func (dnsconf *dNSConfig) GetNameServers() (addrs []string, err error) {
	addrs, err = resolvconf.GetNameServers(dnsconf.resolvConfigPath)
	return addrs, err
}

func (dnsconf *dNSConfig) AddNameServers(addrs []string) (err error) {
	err = dhclient.AddNameServers(addrs, dnsconf.dhclientConfigPath, dnsconf.dhclientConfigPathBackup)
	if err != nil {
		return err
	}

	err = dnsconf.ReloadNameServers()

	return err
}

func (dnsconf *dNSConfig) DHCPNameServers() (err error) {
	err = dhclient.RemoveNameServers(dnsconf.dhclientConfigPath, dnsconf.dhclientConfigPathBackup)
	if err != nil {
		return err
	}

	err = dnsconf.ReloadNameServers()

	return err
}
