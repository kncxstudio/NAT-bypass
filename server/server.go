package server

import (
	"log"
	. "natbypass/utils"
	"net"
	"strconv"
)

func StartServer() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:1199")
	CheckErr(err)
	conn, err := net.ListenUDP("udp", addr)
	CheckErr(err)
	defer conn.Close()
	hostA := ""
	hostB := ""
	for {
		data := make([]byte, 1024)
		_, srcIP, err := conn.ReadFromUDP(data)
		CheckErr(err)
		log.Println("srr ip port:", srcIP.Port)
		// peerAddr := conn.RemoteAddr().String()
		if "" == hostA {
			hostA = srcIP.IP.String() + ":" + strconv.Itoa(srcIP.Port)
			log.Println("A source IP:", hostA)
		} else if "" == hostB {
			hostB = srcIP.IP.String() + ":" + strconv.Itoa(srcIP.Port)
			log.Println("B source IP:", hostB)

			addrA, err := net.ResolveUDPAddr("udp", hostA)
			CheckErr(err)
			conn.WriteToUDP([]byte(hostB), addrA)
			addrB, err := net.ResolveUDPAddr("udp", hostB)
			CheckErr(err)
			conn.WriteToUDP([]byte(hostA), addrB)

			log.Println("IP data send finish!")

			hostA = ""
			hostB = ""
		}

	}
}
