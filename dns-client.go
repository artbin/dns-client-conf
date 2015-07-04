/*
Package dnsclientconf provides a portable interface for operating system
dns client settings.
*/
package dnsclientconf

import "net"

// DNSConfigurator is the interface that wraps the NameServers methods. In
// most cases all methods need Administrator priviledges(except GetNameServers).
//
// GetNameServers is the method that returns the list of host dns addresses.
// Side effect free.
//
// AddNameServers reads the list of ip addresses and changes host
// operating system dns settings. If you want cancel this method effect,
// please call DHCPNameServers.
//
// DHCPNameServers is the method that revert previously configured name
// servers addresses. Generally applies dhcp protocol.
//
// ReloadNameServers is the method that safety refresh dns settings.
type DNSConfigurator interface {
	GetNameServers() ([]net.IP, error)
	AddNameServers(addrs []net.IP) error
	DHCPNameServers() error
	ReloadNameServers() error
	SetInterface(iface *net.Interface)
}
