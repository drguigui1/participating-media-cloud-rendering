package main

import (
    //"fmt"

    //"volumetric-cloud/vector3"
    //"volumetric-cloud/ray"
    "volumetric-cloud/render"
    //"volumetric-cloud/camera"
)

func main() {
    imgSizeX := 256
    imgSizeY := 256
    image := render.Render(imgSizeY, imgSizeX)
    image.SavePPM("tmp.ppm")

}
