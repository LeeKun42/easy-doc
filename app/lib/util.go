package lib

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() []string {
	var ips []string
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface ipAddress：%v\n", err)
		return ips
	}
	for _, address := range interfaceAddr {
		ipNet, isVailIpNet := address.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
