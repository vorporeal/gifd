// http.go
package main

import (
	"image/gif"
	"log"
	"net/http"
)

type HttpAnimServer struct {
	AnimServer
}

func (server HttpAnimServer) Serve(anim *Anim, addrspec string) {
	handler := func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "image/gif")
		err := gif.EncodeAll(res, anim.EncodeGIF())
		if err != nil {
			log.Print(err)
		}
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(addrspec, nil)
}
