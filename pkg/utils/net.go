package utils

import (
	"errors"
	"net"
)

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("Get Local IP error.")
}
