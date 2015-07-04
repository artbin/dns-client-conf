package dnsclientconf

import (
	"net"

	"github.com/ArtemKulyabin/dns-client-conf/debugmode"
	"github.com/ArtemKulyabin/dns-client-conf/resolvconf"
)

const ResolvConfigPath = "/etc/resolv.conf"

type dNSConfig struct {
	resolvConfigPath, Interface string
}

func NewDNSConfigurator() DNSConfigurator {
	return &dNSConfig{ResolvConfigPath, "en0"}
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
	err = debugmode.DebugExec("ipconfig", "set", dnsconf.Interface, "DHCP")
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
