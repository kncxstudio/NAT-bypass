package main

import (
	"natbypass/client"
	"natbypass/server"
	"os"
)

func main() {
	args := os.Args
	if "server" == args[1] {
		server.StartServer()
	} else if "client" == args[1] {
		client.StartClient(args[2], args[3])
	}
}
