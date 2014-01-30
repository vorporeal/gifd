// flags.go
package main

import (
	"flag"
	"log"
	"os"
)

type flagSet struct {
	Http bool
	Path string
	Port uint
}

var flags flagSet

func (f *flagSet) Init() {
	flag.UintVar(&f.Port, "port", 6926,
		"The port on which to listen for requests.")
	flag.StringVar(&f.Path, "path", "",
		"The path to the GIF which will be served.")
	flag.BoolVar(&f.Http, "http", false,
		"Whether the GIF should be served over http.")

	flag.Parse()

	if len(f.Path) == 0 {
		f.ParseError("No path specified.")
	}

	if f.Port >= 1<<16 {
		f.ParseError("Port out of bounds.")
	}
}

func (f *flagSet) ParseError(msg string) {
	log.Println("Error parsing flags:", msg)
	flag.Usage()
	os.Exit(1)
}
