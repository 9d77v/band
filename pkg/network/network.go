package network

import (
	"fmt"
	"log"
	"net"
)

func GetNetworkIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Panicln("net.Interfaces failed, err:", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func GetRandomPort() uint64 {
	address, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	return uint64(listener.Addr().(*net.TCPAddr).Port)
}
