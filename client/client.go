package client

import (
	"fmt"
	"log"
	. "natbypass/utils"
	"net"
	"time"
)

func StartClient(clientID, serverAddrStr string) {
	// svrAddr, err := net.ResolveUDPAddr("udp", serverAddrStr)
	// CheckErr(err)
	conn, err := net.Dial("udp", serverAddrStr)
	CheckErr(err)
	_, err = conn.Write([]byte("first msg" + clientID))
	CheckErr(err)

	data := make([]byte, 1024)
	_, err = conn.Read(data)
	CheckErr(err)
	peerAddr := string(data)
	log.Println("Peer Addr:", peerAddr)
	peerConn, err := net.Dial("udp", peerAddr)
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
		peerConn.Read(peerData)
		fmt.Println(string(peerData))
	}

}
