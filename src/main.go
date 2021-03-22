package main

import (
    //"fmt"

    //"volumetric-cloud/vector3"
    //"volumetric-cloud/ray"
    "volumetric-cloud/img"
    //"volumetric-cloud/camera"
)

func main() {

    res := img.InitImg(256, 256)
    for i := 0; i < 256; i += 1 {
        for j := 0; j < 256; j += 1 {
            // set the pixel value
            res.SetPixel(i, j, 200, 100, 0)
        }
    }

    res.SavePPM("tmp.ppm")

}
