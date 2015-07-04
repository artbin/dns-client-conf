// +build linux freebsd openbsd netbsd

package dnsclientconf

import (
	"net"

	"github.com/ArtemKulyabin/dns-client-conf/dhclient"
	"github.com/ArtemKulyabin/dns-client-conf/resolvconf"
)

type dNSConfig struct {
	resolvConfigPath, dhclientConfigPath, dhclientConfigPathBackup string
	iface                                                          *net.Interface
}

func NewDNSConfigurator() DNSConfigurator {
	iface, _ := net.InterfaceByName(InterfaceName)
	return &dNSConfig{ResolvConfigPath, DhclientConfigPath, DhclientConfigPathBackup, iface}
}

func (dnsconf *dNSConfig) GetNameServers() (addrs []net.IP, err error) {
	return resolvconf.GetNameServers(dnsconf.resolvConfigPath)
}

func (dnsconf *dNSConfig) AddNameServers(addrs []net.IP) (err error) {
	err = dhclient.AddNameServers(addrs, dnsconf.dhclientConfigPath, dnsconf.dhclientConfigPathBackup)
	if err != nil {
		return err
	}

	return dnsconf.ReloadNameServers()
}

func (dnsconf *dNSConfig) DHCPNameServers() (err error) {
	err = dhclient.RemoveNameServers(dnsconf.dhclientConfigPath, dnsconf.dhclientConfigPathBackup)
	if err != nil {
		return err
	}

	return dnsconf.ReloadNameServers()
}

func (dnsconf *dNSConfig) SetInterface(iface *net.Interface) {
	dnsconf.iface = iface
}
