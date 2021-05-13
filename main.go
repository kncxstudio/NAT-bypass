package main

import (
	"flag"
	"natbypass/client"
	"natbypass/server"
	"os"
)

func main() {
	args := os.Args
	if "server" == args[1] {
		server.StartServer()
	} else if "client" == args[1] {
		clientID := flag.String("id", "testClient", "client ID")
		serverAddrStr := flag.String("server", "127.0.0.1:1199", "server address")

		client.StartClient(*clientID, *serverAddrStr)
	}
}
