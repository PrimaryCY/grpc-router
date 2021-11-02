package tools

import "net"

// 获取本地ip地址
func LocalIP() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addr {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
