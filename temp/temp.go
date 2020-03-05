package main

import (
	"net"
	"fmt"
)

func main() {
	localIP, remoteNet, err :=net.ParseCIDR("192.168.9.138/21")
	if err != nil {

	}
	fmt.Println(localIP,remoteNet.IP,remoteNet.Mask)
}
