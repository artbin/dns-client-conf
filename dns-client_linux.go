package dnsclientconf

import "os/exec"

const (
	ResolvConfigPath         = "/etc/resolv.conf"
	DhclientConfigPath       = "/etc/dhcp/dhclient.conf"
	DhclientConfigPathBackup = "/etc/dhcp/dhclient.conf.auto"
)

// For details please see http://www.cyberciti.biz/faq/howto-linux-renew-dhcp-client-ip-address/
func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	dhclientCmd := exec.Command("dhclient", "-r")
	err = dhclientCmd.Run()
	if err != nil {
		return err
	}

	dhclientCmd = exec.Command("dhclient")
	err = dhclientCmd.Run()
	if err != nil {
		return err
	}

	networkManagerCmd := exec.Command("service", "network-manager", "restart")
	err = networkManagerCmd.Run()

	return err
}
