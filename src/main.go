package main

import (
    "math"
    "fmt"

    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/scene"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/camera"
    "volumetric-cloud/vector3"
)

func main() {
    imgSizeX := 600
    imgSizeY := 500

    // Camera
    aspectRatio := float64(imgSizeX) / float64(imgSizeY)
    fieldOfView := math.Pi / 2
    origin := vector3.InitVector3(2.0, 3.0, 3.0)
    camera := camera.InitCamera(
       aspectRatio,
       fieldOfView,
       imgSizeX,
       imgSizeY,
       origin,
       -math.Pi / 8.0,
       0.0,
       0.0,
    )

    // Voxel Grid
    shift := vector3.InitVector3(-1.0, 0.0, -2.0)
    oppositeCorner := vector3.InitVector3(3.0, 2.0, -6.0)
    voxelGrid := voxel_grid.InitVoxelGrid(1.0, shift, oppositeCorner)

    // Spheres
    sphere := sphere.InitSphere(vector3.InitVector3(0, 0, -2), 1.0)

    // Lights
    light := light.InitLight(vector3.InitVector3(0.0, 5.0, 0.0), vector3.InitVector3(100.0 / 255.0, 100.0 / 255.0, 100.0 / 255.0))

    // Scene
    step := 0.2 // step of the ray marching
    s := scene.InitScene(voxelGrid, sphere, camera, light, step);
    fmt.Println(s.VoxelGrid)

    // Render
    image := s.Render(imgSizeY, imgSizeX)

    // Save
    image.SavePPM("tmp.ppm")
}
