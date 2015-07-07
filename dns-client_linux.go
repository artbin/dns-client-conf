package dnsclientconf

import "github.com/ArtemKulyabin/dns-client-conf/debugmode"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhcp/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhcp/dhclient.conf.auto"
	InterfaceName            = ""
)

// For details please see http://www.cyberciti.biz/faq/howto-linux-renew-dhcp-client-ip-address/
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	err = debugmode.DebugExec("dhclient", "-r")
	if err != nil {
		return err
	}

	err = debugmode.DebugExec("dhclient")
	if err != nil {
		return err
	}

	debugmode.DebugExec("service", "network-manager", "restart")

	return err
}
