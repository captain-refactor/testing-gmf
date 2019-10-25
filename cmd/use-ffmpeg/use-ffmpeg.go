package main

import "github.com/3d0c/gmf"

func main() {
	ctx, err := gmf.NewInputCtx("video.mkv")
	defer ctx.Close()
	if err != nil {
		panic(err)
	}
	println()
	for i := 0; i < ctx.StreamsCnt(); i++ {
		stream, err := ctx.GetStream(i)
		if err != nil {
			panic(err)
		}
		println(stream.IsAudio())
	}
}
