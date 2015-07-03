package dnsclientconf

import "testing"

func TestAddNameServers(t *testing.T) {

	dnsconf := NewDNSConfigurator()

	err := dnsconf.AddNameServers([]string{"8.8.8.8", "8.8.4.4"})
	if err != nil {
		t.Error("AddNameServers ", err)
	}
}

func TestGetNameServers(t *testing.T) {

	dnsconf := NewDNSConfigurator()

	_, err := dnsconf.GetNameServers()

	if err != nil {
		t.Error("GetNameServers ", err)
	}
}

func TestDHCPNameServers(t *testing.T) {

	dnsconf := NewDNSConfigurator()

	err := dnsconf.DHCPNameServers()

	if err != nil {
		t.Error("DHCPNameServers ", err)
	}
}

func TestReloadNameServers(t *testing.T) {

	dnsconf := NewDNSConfigurator()

	err := dnsconf.ReloadNameServers()

	if err != nil {
		t.Error("ReloadNameServers ", err)
	}
}
