package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Pass the port as the first argument")
	}
	// A map of key to address
	addrMap := make(map[string]*net.UDPAddr)
	// Listen
	addr := net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
	}
	addr.Port, _ = strconv.Atoi(os.Args[1])
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalln("cannot listen:", err)
	}
	defer conn.Close()
	log.Println("Listening on", conn.LocalAddr())
	// Wait for connections
	buffer := make([]byte, 128)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error: ", err)
		}
		key := string(buffer[:n])
		// Check what to do
		if otherAddress, exists := addrMap[key]; exists {
			delete(addrMap, key)
			log.Println("Punching", addr.String(), "with", otherAddress.String(), "with key", key)
			_, _ = conn.WriteTo([]byte(addr.String()), otherAddress)
			_, _ = conn.WriteTo([]byte(otherAddress.String()), addr)
			log.Println("Punched", addr.String(), "with", otherAddress.String(), "with key", key)
		} else {
			// Create a key
			addrMap[key] = addr
			log.Println("Registered", addr.String(), "with", key)
		}
		log.Println("Read", string(buffer[:n]), "from", addr.String())
	}
}
