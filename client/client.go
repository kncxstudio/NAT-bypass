package client

import (
	"fmt"
	"log"
	. "natbypass/utils"
	"net"
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

	peerConn, err := net.Dial("udp", peerAddr)
	CheckErr(err)
	go func() {
		for {
			_, err := peerConn.Write([]byte("msg from " + clientID))
			if CheckErr(err) {
				log.Println("send msg successfully")
			}
		}
	}()

	for {
		peerData := make([]byte, 1024)
		peerConn.Read(peerData)
		fmt.Println(string(peerData))
	}

}
