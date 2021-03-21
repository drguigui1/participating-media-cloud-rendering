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

    image := img.Img{0, 0, []vector3.Vector3{}}
    fmt.Println(image)

    res := img.InitImg(10, 10)
    fmt.Println(res)
}
