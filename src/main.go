package main

import (
    "fmt"

    "volumetrical-cloud/vector3"
)

func main() {
    fmt.Println(vector3.InitVector3(1.0 / 14.0, 2.0 / 14.0, 3.0 / 14.0).Length())
}
