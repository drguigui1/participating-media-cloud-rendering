package img

import (
    "volumetric-cloud/vector3"
)

type Img struct {
    NbRows int
    NbCols int
    Pixels []vector3.Vector3
}

func (img *Img) SetPixel(i, j int, color vector3.Vector3) {
    img.Pixels[i * img.NbCols + j] = color
}

func (img Img) GetPixel(i, j int) vector3.Vector3 {
    return img.Pixels[i * img.NbCols + j]
}
