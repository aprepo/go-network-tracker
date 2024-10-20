package main

import (
	"fmt"
	"log"
	"strings"

	gopsutil "github.com/shirou/gopsutil/v3/net"
)

func isLocalAddr(addr string) bool {
	privateRanges := []string{
		"10.",
		"127.",
		"172.",
		"192.168.",
	}

	for _, privateRange := range privateRanges {
		if strings.HasPrefix(addr, privateRange) {
			return true
		}
	}
	return false
}

func main() {
	stats, err := gopsutil.IOCounters(true)
	if err != nil {
		log.Fatal("Failed to get IOCounters")
		return
	}
	for _, stat := range stats {
		fmt.Printf("Interface: %s Sent: %d Recv: %d\n", stat.Name, stat.BytesSent, stat.BytesRecv)
	}

	conns, err := gopsutil.Connections("all")
	if err != nil {
		log.Fatal("Error getting connections")
	}
	for _, con := range conns {
		var localIP string = con.Laddr.IP
		var remoteIP string = con.Raddr.IP
		if !isLocalAddr(remoteIP) {
			fmt.Printf("Connected from: %v:%v to: %v:%v STATE: %v\n", localIP, con.Laddr.Port, remoteIP, con.Raddr.Port, con.Status)
		}
	}
}
