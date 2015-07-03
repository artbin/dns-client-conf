package dnsclientconf

import "os/exec"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhclient.conf.auto"
)

// For details please see http://www.cyberciti.biz/faq/freebsd-unix-force-dhcp-client-to-get-a-new-lease/
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	dhclientCmd := exec.Command("service", "dhclient", "restart", "em0")
	err = dhclientCmd.Run()

	return err
}
