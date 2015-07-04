package dnsclientconf

import "github.com/ArtemKulyabin/dns-client-conf/debugmode"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhclient.conf.auto"
	InterfaceName            = ""
)

// For details please see http://www.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man8/dhclient.8
//
// When configuring the interface, dhclient attempts to remove any existing
// addresses, gateway routes that use the interface, and non-permanent arp(8)
// entries. Conversely, if the interface is later manipulated to add or
// delete addresses then dhclient will automatically exit. It thus
// automatically exits whenever a new dhclient is run on the same interface.
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	// On receiving HUP dhclient will restart itself, reading dhclient.conf(5) and
	// obtaining a new lease.
	return debugmode.DebugExec("pkill", "-HUP", "dhclient")
}
