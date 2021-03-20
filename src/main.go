package main

import (
    "fmt"

    "volumetrical-cloud/vector3"
    "volumetrical-cloud/ray"
    "volumetrical-cloud/img"
)

func main() {
    ray := ray.InitRay(vector3.InitVector3(1, 1, 1), vector3.InitVector3(2, 2, 2))
    fmt.Println(ray)
    fmt.Println(vector3.InitVector3(1.0 / 14.0, 2.0 / 14.0, 3.0 / 14.0).Length())

    img := img.Img{0, 0, []vector3.Vector3{}}
    fmt.Println(img)
}
