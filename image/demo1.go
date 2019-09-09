package main

import (
    "image"
    "math"
    "image/color"
    "image/gif"
    "io"
    "math/rand"
    "os"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

func main() {
    file, _ := os.Create("./image/math2.gif")
    lissajous(file)
}
func lissajous(out io.Writer) {
    const (
        cycles  = 5 //
        res     = 0.001
        size    = 100
        nframes = 64 // 设置gif的帧数
        delay   = 8
    )
    freq := rand.Float64() * 3.0 //
    // 创建gif动画数组
    anim := gif.GIF{LoopCount: nframes}
    //
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            // 给图片指定的位置填充颜色
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
        }
        phase += 0.1
        // 设置每一帧的延时
        anim.Delay = append(anim.Delay, delay)
        // 设置每一帧的图片
        anim.Image = append(anim.Image, img)
    }
    // 将编码后的gif动画文件 保存至指定的文件中
    gif.EncodeAll(out, &anim)
}