// gifd.go
package main

import (
	"fmt"
	"log"
	"os"
)

type AnimServer interface {
	Serve(*Anim, string)
}

func main() {
	flags.Init()

	// Test the anim code.
	f, err := os.Open(flags.Path)
	if err != nil {
		log.Fatal(err)
	}
	anim := NewAnim(f)

	var server AnimServer
	if flags.Http {
		server = HttpAnimServer{}
	} else {
		server = TelnetAnimServer{}
	}

	addrspec := fmt.Sprintf("localhost:%d", flags.Port)
	log.Printf("Serving %s at %s...\n",
		flags.Path, addrspec)
	server.Serve(anim, addrspec)
}
