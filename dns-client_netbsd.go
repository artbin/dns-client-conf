package dnsclientconf

import "github.com/ArtemKulyabin/dns-client-conf/debugmode"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhclient.conf.auto"
)

// For details please see https://www.netbsd.org/docs/network/dhcp.html#enable-dhcp
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	return debugmode.DebugExec("sh", "/etc/rc.d/dhclient", "restart")
}
