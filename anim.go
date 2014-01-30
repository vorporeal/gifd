// anim.go
package main

import (
	"image"
	"image/gif"
	"io"
	"log"
)

type Anim struct {
	raw        *gif.GIF
	frames     []image.Image
	frameCount int
	duration   int
	elapsed    []int
	// TODO(vorporeal): Keep track of whether the GIF should loop.
}

func NewAnim(r io.Reader) *Anim {
	a := new(Anim)
	g, err := gif.DecodeAll(r)
	if err != nil {
		log.Fatal(err)
	}
	a.raw = g
	a.frameCount = len(g.Image)
	log.Printf("Loaded GIF containing %d frames.\n",
		a.frameCount)

	a.frames = make([]image.Image, a.frameCount)
	a.elapsed = make([]int, a.frameCount)
	for i := 0; i < a.frameCount; i++ {
		// Generate an Image for each frame
		p := g.Image[i]
		a.frames[i] = p.SubImage(p.Rect)

		// The frame delay is specified in 100ths of a second,
		// so multiply by 10 to get milliseconds.
		frameLength := g.Delay[i] * 10
		frameStart := 0
		if i > 0 {
			frameStart = a.elapsed[i-1] + frameLength
		}
		a.elapsed[i] = frameStart
		a.duration += frameLength
	}
	log.Printf("Total duration: %.2fs\n", float32(a.duration)/1000.0)
	return a
}

func (a *Anim) ImageAtTime(t int) image.Image {
	for elapsed, i := range a.elapsed {
		if elapsed <= t {
			log.Printf("frame %d is at time %d", i, t)
			return a.frames[i]
		}
	}
	return nil
}

func (a *Anim) EncodeGIF() *gif.GIF {
	// TODO(vorporeal): Actually implement this method, so Anim doesn't need to
	// 		contain a GIF.
	return a.raw
}
