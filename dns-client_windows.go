package dnsclientconf

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ArtemKulyabin/dns-client-conf/helpers"
)

type dNSConfig struct {
	Interface string
}

func NewDNSConfigurator() DNSConfigurator {
	return &dNSConfig{"Local Area Connection"}
}

// For details please see http://stackoverflow.com/questions/18620173/how-can-i-set-change-dns-using-the-command-prompt-at-windows-8
// and https://technet.microsoft.com/en-us/library/cc778503(v=ws.10).aspx

func (dnsconf *dNSConfig) AddNameServers(addrs []string) (err error) {
	for _, addr := range addrs {
		err = helpers.CheckIP(addr)
		if err != nil {
			return err
		}
	}

	for i, addr := range addrs {
		netshCmd := exec.Command("netsh", "interface", "ip", "add", "dnsservers", dnsconf.Interface, addr, fmt.Sprintf("%d", i+1))
		err = netshCmd.Run()
		if err != nil {
			return err
		}
	}

	err = dnsconf.ReloadNameServers()

	return err
}

func (dnsconf *dNSConfig) DHCPNameServers() (err error) {
	netshCmd := exec.Command("netsh", "interface", "ip", "set", "dnsservers", dnsconf.Interface, "dhcp")
	err = netshCmd.Run()
	if err != nil {
		return err
	}

	err = dnsconf.ReloadNameServers()

	return err
}

func (dnsconf *dNSConfig) ReloadNameServers() (err error) {
	ipconfigCmd := exec.Command("ipconfig", "/flushdns")
	err = ipconfigCmd.Run()

	return err
}

func (dnsconf *dNSConfig) GetNameServers() (addrs []string, err error) {
	netshCmd := exec.Command("netsh", "interface", "ip", "show", "dnsservers", dnsconf.Interface)
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
				addrs = append(addrs, addr)
			}
			err = nil
		}
	}

	return addrs, err
}
