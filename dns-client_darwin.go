package dnsclientconf

import (
	"os/exec"

	"github.com/ArtemKulyabin/dns-client-conf/resolvconf"
)

const ResolvConfigPath = "/etc/resolv.conf"

type dNSConfig struct {
	resolvConfigPath, Interface string
}

func NewDNSConfigurator() DNSConfigurator {
	return &dNSConfig{ResolvConfigPath, "en0"}
}

func (dnsconf *dNSConfig) AddNameServers(addrs []string) (err error) {
	//networksetup -setdnsservers Ethernet addrs
	args := []string{"-setdnsservers", "Ethernet"}
	args = append(args, addrs...)
	networksetupCmd := exec.Command("networksetup", args...)
	err = networksetupCmd.Run()
	if err != nil {
		return err
	}

	err = dnsconf.ReloadNameServers()

	return err
}

func (dnsconf *dNSConfig) GetNameServers() (addrs []string, err error) {
	addrs, err = resolvconf.GetNameServers(dnsconf.resolvConfigPath)
	return addrs, err
}

func (dnsconf *dNSConfig) DHCPNameServers() (err error) {
	ipconfigCmd := exec.Command("ipconfig", "set", dnsconf.Interface, "DHCP")
	err = ipconfigCmd.Run()
	if err != nil {
		return err
	}

	err = dnsconf.ReloadNameServers()

	return err
}

func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	discoveryutilCmd := exec.Command("discoveryutil", "mdnsflushcache")
	err = discoveryutilCmd.Run()
	if err != nil {
		return err
	}

	discoveryutilCmd = exec.Command("discoveryutil", "udnsflushcaches")
	err = discoveryutilCmd.Run()

	return err
}
