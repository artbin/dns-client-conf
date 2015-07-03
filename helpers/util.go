package helpers

import (
	"errors"
	"fmt"
	"net"
)

func CheckIP(addr string) (err error) {
	if net.ParseIP(addr) == nil {
		err = errors.New(fmt.Sprintf("IP address %s is invalid", addr))
		return err
	}

	return err
}
