package resolvconf

import (
	"bufio"
	"net"
	"os"
	"strings"

	"github.com/ArtemKulyabin/dns-client-conf/helpers"
)

// GetNameServers consume resolv.conf path and returns current host dns servers
// addresses.
func GetNameServers(filename string) (addrs []net.IP, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
		switch lineArr[0] {
		case "nameserver": // add one name server
			if len(lineArr) > 1 {
				err = helpers.CheckIP(lineArr[1])
				if err != nil {
					return addrs, err
				}
				addrs = append(addrs, net.ParseIP(lineArr[1]))
			}

		}
	}

	return addrs, err
}
