package main

import (
    "fmt"
    "math"

    "volumetric-cloud/vector3"
    "volumetric-cloud/ray"
    "volumetric-cloud/img"
    "volumetric-cloud/camera"
)

func main() {
    ray := ray.InitRay(vector3.InitVector3(1, 1, 1), vector3.InitVector3(2, 2, 2))
    fmt.Println(ray)
    fmt.Println(vector3.InitVector3(1.0 / 14.0, 2.0 / 14.0, 3.0 / 14.0).Length())

    img := img.Img{0, 0, []vector3.Vector3{}}
    fmt.Println(img)

    fmt.Println("----------")

    var imgW int = 600
    var imgH int = 500
    aspectRatio := float64(imgW) / float64(imgH)
    fieldOfView := math.Pi / 2.0
    origin := vector3.InitVector3(0.0, 2.0, 0.0)
    newCamera := camera.InitCamera(
        aspectRatio,
        fieldOfView,
        imgW,
        imgH,
        origin,
        -math.Pi / 4.0,
        0.0,
        0.0,
    )

    fmt.Println(newCamera)
}
