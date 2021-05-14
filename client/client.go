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
	_, err = conn.Read(data)
	CheckErr(err)
	peerAddrStr := string(data)
	log.Println("Peer Addr:", peerAddrStr)
	peerAddr, _ := net.ResolveUDPAddr("udp", peerAddrStr)
	peerConn, err := net.DialUDP("udp", srcAddr, peerAddr)
	CheckErr(err)
	go func() {
		for {
			_, err := peerConn.Write([]byte("msg from " + clientID))
			if CheckErr(err) {
				log.Println("send msg successfully")
			}
			time.Sleep(1)
		}
	}()

	for {
		peerData := make([]byte, 1024)
		_, _, err := peerConn.ReadFromUDP(peerData)
		CheckErr(err)
		fmt.Println(string(peerData))
	}

}
