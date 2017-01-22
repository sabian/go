package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

var palette = []color.Color{
    color.Black,
    color.White,
    color.RGBA{0, 102, 0, 1},
    color.RGBA{0, 86, 255, 1},
    color.RGBA{255, 192, 0, 1},
    color.RGBA{95, 186, 125, 1},
    color.RGBA{254, 228, 197, 1},
    color.RGBA{239, 240, 241, 1},
    color.RGBA{244, 128, 36, 1},
    color.RGBA{177, 10, 7, 1}}

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5
        res = 0.001
        size = 100
        nframes = 64
        delay = 8
    )

    rand.Seed(time.Now().UTC().UnixNano())
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0

    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)

        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            colorIndex := uint8(rand.Intn(9))

            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }

    gif.EncodeAll(out, &anim)
}
