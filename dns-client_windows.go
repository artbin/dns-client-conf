package dnsclientconf

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/ArtemKulyabin/dns-client-conf/debugmode"
	"github.com/ArtemKulyabin/dns-client-conf/helpers"
)

const InterfaceName = "Local Area Connection"

type dNSConfig struct {
	iface *net.Interface
}

func NewDNSConfigurator() DNSConfigurator {
	iface, _ := net.InterfaceByName(InterfaceName)
	return &dNSConfig{iface}
}

// For details please see http://stackoverflow.com/questions/18620173/how-can-i-set-change-dns-using-the-command-prompt-at-windows-8
// and https://technet.microsoft.com/en-us/library/cc778503(v=ws.10).aspx

func (dnsconf *dNSConfig) AddNameServers(addrs []net.IP) (err error) {
	for i, addr := range addrs {
		err = debugmode.DebugExec("netsh", "interface", "ip", "add", "dnsservers", dnsconf.iface.Name, addr.String(), fmt.Sprintf("%d", i+1))
		if err != nil {
			return err
		}
	}

	return dnsconf.ReloadNameServers()
}

func (dnsconf *dNSConfig) DHCPNameServers() (err error) {
	err = debugmode.DebugExec("netsh", "interface", "ip", "set", "dnsservers", dnsconf.iface.Name, "dhcp")
	if err != nil {
		return err
	}

	return dnsconf.ReloadNameServers()
}

func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	return debugmode.DebugExec("ipconfig", "/flushdns")
}

func (dnsconf *dNSConfig) GetNameServers() (addrs []net.IP, err error) {
	netshCmd := exec.Command("netsh", "interface", "ip", "show", "dnsservers", dnsconf.iface.Name)
	outputbuf, err := netshCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	output := bytes.NewBuffer(outputbuf).String()

	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		scannerLine := bufio.NewScanner(strings.NewReader(line))
		scannerLine.Split(bufio.ScanWords)
		var lineArr []string
		for scannerLine.Scan() {
			lineArr = append(lineArr, scannerLine.Text())
		}

		//empty line
		if len(lineArr) == 0 {
			continue
		}
		for _, addr := range lineArr {
			err = helpers.CheckIP(addr)
			if err == nil {
				addrs = append(addrs, net.ParseIP(addr))
			}
			err = nil
		}
	}

	return addrs, err
}

func (dnsconf *dNSConfig) SetInterface(iface *net.Interface) {
	dnsconf.iface = iface
}
