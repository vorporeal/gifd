// telnet.go
package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/ziutek/telnet"
	"image"
	"image/color"
	"io"
	"log"
	"math"
	"net"
	"strings"
)

const (
	ESC   = string('\x1B')
	SCALE = 1
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
		log.Print(conn.RemoteAddr(), " ] Client connected.")

		// Wrap the connection as a telnet connection
		tc, err := telnet.NewConn(conn)
		if err != nil {
			log.Print("Error opening connection:", err)
			continue
		}

		go serveGIF(tc, anim)
	}
}

func serveGIF(conn *telnet.Conn, anim *Anim) {
	writeImage(conn, anim.ImageAtTime(500))

	// Close the connection when done serving
	err := conn.Close()
	if err != nil {
		log.Print("Error closing connection:", err)
		return
	}
	log.Print(conn.RemoteAddr(), " ] Client disconnected.")
}

func writeImage(w io.Writer, img image.Image) {
	rect := img.Bounds()
	width := uint(SCALE * float32(rect.Max.X-rect.Min.X))
	height := uint(SCALE * float32(rect.Max.Y-rect.Min.Y))
	scaled := resize.Resize(width, 0, img, resize.NearestNeighbor)

	for y := 0; y < int(height); y++ {
		line := ""
		for x := 0; x < int(width); x++ {
			line += terminalCodeForColor(scaled.At(x, y)) + "  "
		}
		line += fmt.Sprintf("%s[0m\n", ESC)
		io.Copy(w, strings.NewReader(line))
	}
}

func terminalCodeForColor(color color.Color) string {
	r, g, b, _ := color.RGBA()
	code := 16
	if r == g && g == b {
		if r != 0 {
			increment := int(math.Ceil(float64(r) / math.MaxUint16 * 24))
			code = 232 + increment
			// TODO(vorporeal): Fix white not working properly. Write tests?
		}
	} else {
		convert := func(val uint32) int {
			return int(math.Ceil(float64(val) / math.MaxUint16 * 5))
		}
		code += 36*convert(r) + 6*convert(g) + convert(b)
	}

	return fmt.Sprintf("%s[48;5;%dm", ESC, code)
}
