package dnsclientconf

import "github.com/ArtemKulyabin/dns-client-conf/debugmode"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhclient.conf.auto"
	InterfaceName            = "em0"
)

// For details please see http://www.cyberciti.biz/faq/freebsd-unix-force-dhcp-client-to-get-a-new-lease/
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	return debugmode.DebugExec("service", "dhclient", "restart", "em0")
}
