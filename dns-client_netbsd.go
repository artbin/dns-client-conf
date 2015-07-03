package dnsclientconf

import "os/exec"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhclient.conf.auto"
)

// For details please see https://www.netbsd.org/docs/network/dhcp.html#enable-dhcp
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	dhclientCmd := exec.Command("sh", "/etc/rc.d/dhclient", "restart")
	err = dhclientCmd.Run()
	if err != nil {
		return err
	}

	return err
}
