package utils

import (
	"fmt"
	"net"
	"time"
)


func CheckPorts(host string, ports []string) {
	for _, port := range ports {
		timeout:= time.Second
		conn, err:= net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Printf("Port %s is closed,\n", port)
		}
		if conn != nil {
			fmt.Printf("Port %s is open,\n", port)
			defer conn.Close()
		}
	}
	
}