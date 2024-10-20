package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Test")
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Failed to get network interfaces", err)
		return
	}

	for _, iface := range interfaces {
		fmt.Printf("Interface: %v\n", iface.Name)
		addresses, err := iface.Addrs()
		if err != nil {
			log.Fatal("Failed to get addresses", err)
			continue
		}
		for _, addr := range addresses {
			fmt.Printf(" Address: %v\n", addr.String())
		}
	}
}
