package img

import (
    "os"
    "fmt"

    "volumetric-cloud/vector3"
)

type Img struct {
    NbRows int
    NbCols int
    Pixels []vector3.Vector3
}

func InitImg(nbRows, nbCols int) Img {
    return Img{
        NbRows: nbRows,
        NbCols: nbCols,
        Pixels: make([]vector3.Vector3, nbRows * nbCols),
    }
}

func (img *Img) SetPixel(i, j int, color vector3.Vector3) {
    img.Pixels[i * img.NbCols + j] = color
}

func (img Img) GetPixel(i, j int) vector3.Vector3 {
    return img.Pixels[i * img.NbCols + j]
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
    content := make([]byte, len(img.Pixels) * 3)
    j := 0
    for i := 0; i < len(img.Pixels); i += 1 {
        v := vector3.MulVector3(img.Pixels[i], 255.0)
        content[j] = byte(v.X)
        content[j + 1] = byte(v.Y)
        content[j + 2] = byte(v.Z)
        j += 3
    }

    _, err = file.Write(content)
    if err != nil {
        return err
    }

    return nil
}
