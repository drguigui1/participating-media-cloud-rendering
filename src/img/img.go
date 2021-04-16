package img

import (
    "os"
    "image"
    "image/color"
    "image/png"
)

type Img struct {
    Width int
    Height int
    Image *image.NRGBA
}

func InitImg(width, height int) Img {
    return Img{
        Width: width,
        Height: height,
        Image: image.NewNRGBA(image.Rect(0, 0, width, height)),
    }
}

// i -> width
// j -> height
func (img *Img) SetPixel(i, j int, r, g, b, a uint8) {
    img.Image.Set(i, j, color.NRGBA{
        R: r,
        G: g,
        B: b,
        A: a,
    })
}

func (img Img) SavePNG(filepath string) {
    f, err := os.Create(filepath)
    if err != nil {
        panic(err)
    }

    if err := png.Encode(f, img.Image); err != nil {
        f.Close()
        panic(err)
    }

    if err := f.Close(); err != nil {
        panic(err)
    }
}


