package main

import (
    "fmt"

    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/img"
    //"volumetric-cloud/camera"
)

func main() {
    ray := ray.InitRay(vector3.InitVector3(1, 1, 1), vector3.InitVector3(2, 2, 2))
    fmt.Println(ray)
    fmt.Println(vector3.InitVector3(1.0 / 14.0, 2.0 / 14.0, 3.0 / 14.0).Length())

    res := img.InitImg(256, 256)
    for i := 0; i < 256 * 256; i += 1 {
        res.Pixels[i] = vector3.InitVector3(1.0, 0.2, 0.2)
    }

    res.SavePPM("tmp.ppm")

}
