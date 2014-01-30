// telnet.go
package main

import (
	"log"
	"net"
)

type TelnetAnimServer struct {
	AnimServer
}

func (server TelnetAnimServer) Serve(anim *Anim, addrspec string) {
	ln, err := net.Listen("tcp", addrspec)
	if err != nil {
		log.Print("Error opening port:", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("Error opening connection:", err)
			continue
		}
		remoteAddr := conn.RemoteAddr().String()
		log.Print(remoteAddr, " ] Client connected.")

		// TODO(vorporeal): Start a goroutine which serves the GIF.

		err = conn.Close()
		if err != nil {
			log.Print("Error closing connection:", err)
			continue
		}
		log.Print(remoteAddr, " ] Client disconnected.")
	}
}
