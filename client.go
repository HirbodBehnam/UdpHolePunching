package main

import (
	"log"
	"net"
	"os"
	"time"
)

// Periodically send a packet to remote
func punchRemote(conn *net.UDPConn, other string) {
	ip, err := net.ResolveUDPAddr("udp", other)
	if err != nil {
		log.Fatalln("cannot parse other")
	}
	for {
		conn.WriteToUDP([]byte("PUNCHED!"), ip)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Pass the punch server as the first argument and the key as second")
	}
	// Listen on random port
	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Fatalln("cannot listen:", err)
	}
	defer conn.Close()
	log.Println("Initiated a UDP socket at", conn.LocalAddr())
	// Send to puncher server
	puncherAddr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatalln("cannot parse punch server address")
	}
	conn.WriteTo([]byte(os.Args[2]), puncherAddr)
	// Read the result
	buffer := make([]byte, 128)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Println("Error in before punch:", err)
	}
	log.Println("Got", string(buffer[:n]), "from", addr)
	// Punch!
	go punchRemote(conn, string(buffer[:n]))
	for {
		n, addr, err = conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Error:", err)
		}
		log.Println("Got", string(buffer[:n]), "from", addr)
	}
}
