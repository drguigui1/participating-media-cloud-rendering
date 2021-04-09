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

/*type Img struct {
    NbRows int
    NbCols int
    Pixels []byte
}

func InitImg(nbRows, nbCols int) Img {
    return Img{
        NbRows: nbRows,
        NbCols: nbCols,
        Pixels: make([]byte, nbRows * nbCols * 3),
    }
}

func (img *Img) SetPixel(i, j int, r, g, b byte) {
    img.Pixels[i * img.NbCols * 3 + j * 3] = r
    img.Pixels[i * img.NbCols * 3 + j * 3 + 1] = g
    img.Pixels[i * img.NbCols * 3 + j * 3 + 2] = b
}

func (img Img) GetPixel(i, j int) (byte, byte, byte, error) {
    if i < 0 || i > img.NbRows ||
       j < 0 || j > img.NbCols {
           return 0, 0, 0, errors.New("ERROR: Wrong i or j")
    }
    r := img.Pixels[i * img.NbCols * 3 + j * 3]
    g := img.Pixels[i * img.NbCols * 3 + j * 3 + 1]
    b := img.Pixels[i * img.NbCols * 3 + j * 3 + 2]

    return r, g, b, nil
}

func (img Img) SavePPM(path string) error {
    var file *os.File
    var err error

    // create the file
    file, err = os.Create(path)
    if err != nil {
        return err
    }

    // write header
    _, err = fmt.Fprintf(file, "P6\n%d\n%d\n255\n", img.NbCols, img.NbRows)
    if err != nil {
        return err
    }

    // write the content
    _, err = file.Write(img.Pixels)
    if err != nil {
        return err
    }

    return nil
}*/
