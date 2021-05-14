package client

import (
	"fmt"
	"log"
	. "natbypass/utils"
	"net"
	"time"
)

func StartClient(clientID, clientAddrStr, serverAddrStr string) {
	// svrAddr, err := net.ResolveUDPAddr("udp", serverAddrStr)
	// CheckErr(err)
	srcAddr, _ := net.ResolveUDPAddr("udp", clientAddrStr)
	serverAddr, _ := net.ResolveUDPAddr("udp", serverAddrStr)
	conn, err := net.DialUDP("udp", srcAddr, serverAddr)
	CheckErr(err)
	_, err = conn.Write([]byte("first msg" + clientID))
	CheckErr(err)

	data := make([]byte, 1024)
	n, err := conn.Read(data)
	CheckErr(err)
	conn.Close()
	peerAddrStr := string(data[:n])
	log.Println("Peer Addr:", peerAddrStr)
	peerAddr, err := net.ResolveUDPAddr("udp", peerAddrStr)
	CheckErr(err)
	peerConn, err := net.DialUDP("udp", srcAddr, peerAddr)
	CheckErr(err)
	go func() {
		for {
			_, err := peerConn.Write([]byte("msg from " + clientID))
			if CheckErr(err) {
				log.Println("send msg successfully")
			}
			time.Sleep(1000)
		}
	}()

	for {
		peerData := make([]byte, 1024)
		n, _, err := peerConn.ReadFromUDP(peerData)
		CheckErr(err)
		fmt.Println(string(peerData[:n]))
	}

}
