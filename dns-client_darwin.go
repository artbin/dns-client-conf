package dnsclientconf

import (
	"net"

	"github.com/ArtemKulyabin/dns-client-conf/debugmode"
	"github.com/ArtemKulyabin/dns-client-conf/resolvconf"
)

const (
	ResolvConfigPath = "/etc/resolv.conf"
	InterfaceName    = "en0"
)

type dNSConfig struct {
	resolvConfigPath string
	iface            *net.Interface
}

func NewDNSConfigurator() DNSConfigurator {
	iface, _ := net.InterfaceByName(InterfaceName)
	return &dNSConfig{ResolvConfigPath, iface}
}

func (dnsconf *dNSConfig) AddNameServers(addrs []net.IP) (err error) {
	//networksetup -setdnsservers Ethernet addrs
	args := []string{"-setdnsservers", "Ethernet"}
	for _, addr := range addrs {
		args = append(args, addr.String())
	}
	err = debugmode.DebugExec("networksetup", args...)
	if err != nil {
		return err
	}

	return dnsconf.ReloadNameServers()
}

func (dnsconf *dNSConfig) GetNameServers() (addrs []net.IP, err error) {
	return resolvconf.GetNameServers(dnsconf.resolvConfigPath)
}

func (dnsconf *dNSConfig) DHCPNameServers() (err error) {
	err = debugmode.DebugExec("ipconfig", "set", dnsconf.iface.Name, "DHCP")
	if err != nil {
		return err
	}

	return dnsconf.ReloadNameServers()
}

func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	err = debugmode.DebugExec("discoveryutil", "mdnsflushcache")
	if err != nil {
		return err
	}

	return debugmode.DebugExec("discoveryutil", "udnsflushcaches")
}

func (dnsconf *dNSConfig) SetInterface(iface *net.Interface) {
	dnsconf.iface = iface
}
