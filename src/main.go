package main

import (
    "math"

    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/scene"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/camera"
    "volumetric-cloud/vector3"
)

func main() {
    imgSizeX := 1200
    imgSizeY := 1000

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
    shift := vector3.InitVector3(-2.0, 0.0, -6.0)
    oppositeCorner := vector3.InitVector3(5.0, 3.0, -3.0)
    voxelGrid := voxel_grid.InitVoxelGrid(0.1, shift, oppositeCorner, 0.05)

    // Spheres
    sphere := sphere.InitSphere(vector3.InitVector3(0, 0, -2), 1.0)

    // Lights
    //light := light.InitLight(vector3.InitVector3(0.0, 5.0, 0.0), vector3.InitVector3(100.0 / 255.0, 100.0 / 255.0, 100.0 / 255.0))
    light := light.InitLight(vector3.InitVector3(0.0, 5.0, 0.0), vector3.InitVector3(1.0, 1.0, 1.0))

    // Scene
    s := scene.InitScene(voxelGrid, sphere, camera, light);

    // Render
    image := s.Render(imgSizeY, imgSizeX, 3)

    // Save
    image.SavePPM("tmp.ppm")
}
