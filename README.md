# dns-client-conf - Simple and portable cli tool and golang package for work with operating system dns client settings

[![Build Status](https://travis-ci.org/ArtemKulyabin/dns-client-conf.svg)](https://travis-ci.org/ArtemKulyabin/dns-client-conf)

## Installation

To install dns-client-conf, please use `go get`.

### Command line tool

```
$ go get github.com/ArtemKulyabin/dns-client-conf/cmd/dns-client-conf
...
$ dns-client-conf help
...
```

### Package

```
$ go get github.com/ArtemKulyabin/dns-client-conf
...
```

## Usage

### Command line tool

```
$ sudo dns-client-conf add 8.8.8.8 8.8.4.4 # Google Public DNS
DNS servers added successfully
$ dns-client-conf show
8.8.8.8
8.8.4.4
...
$ sudo dns-client-conf add 77.88.8.8 77.88.8.1 # Yandex.DNS
DNS servers added successfully
$ dns-client-conf show
77.88.8.8
77.88.8.1
...
$ sudo dns-client-conf dhcp # Receive dns addresses by DHCP protocol
DNS servers received by DCHP protocol successfully
$ dns-client-conf show # My ISP provider DNS servers
141.105.32.88
141.105.32.89
127.0.1.1
...
```

### Package

```
package main

import (
  "log"
  "github.com/ArtemKulyabin/dns-client-conf"
)

func main() {

  dnsconf := dnsclientconf.NewDNSConfigurator()

  addrs, err := dnsconf.GetNameServers()
  if err != nil {
    log.Fatal(err)
  }

  err = dnsconf.AddNameServers([]string{"8.8.8.8", "8.8.4.4"})
  if err != nil {
    log.Fatal(err)
  }

  err = dnsconf.DHCPNameServers()
  if err != nil {
    log.Fatal(err)
  }
}

```

## Cross compilation
For cross compilation you may use the [Gox](github.com/mitchellh/gox). Example:
```
$ gox github.com/ArtemKulyabin/dns-client-conf/cmd/dns-client-conf
...
```

## Operating system support
* Linux, FreeBSD, OpenBSD, NetBSD, Darwin(OS X), Windows

### Tested on
* Ubuntu 14.04, FreeBSD 10.1, OpenBSD 5.6, NetBSD 6.1.5, OS X Yosemite 10.10.3, Windows 7
