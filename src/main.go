package main

import (
    "math"

    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/scene"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/camera"
    "volumetric-cloud/vector3"
    "volumetric-cloud/noise"
)

func main() {
    imgSizeX := 1200
    imgSizeY := 1000

    // Camera
    aspectRatio := float64(imgSizeX) / float64(imgSizeY)
    fieldOfView := math.Pi / 2
    origin := vector3.InitVector3(-12, 6, -13)
    camera := camera.InitCamera(
       aspectRatio,
       fieldOfView,
       imgSizeX,
       imgSizeY,
       origin,
       -math.Pi / 8,
       -math.Pi / 2,
       0.0,
    )

/*    origin := vector3.InitVector3(0, 3, 5)
    camera := camera.InitCamera(
       aspectRatio,
       fieldOfView,
       imgSizeX,
       imgSizeY,
       origin,
       0.0,
       0.0,
       0.0,
    )*/

    // Voxel Grid 1
    shift := vector3.InitVector3(-4.0, -3.0, -15.0)
    oppositeCorner := vector3.InitVector3(5.0, 1.0, -9.0)
    var seed int64 = 42
    perlinNoise := noise.InitPerlinNoise(1.0, 2.0, 1.0, 0.5, 5, seed)
    voxelGrid := voxel_grid.InitVoxelGrid(0.1, shift, oppositeCorner, 0.05, perlinNoise)

    // Voxel Grid 2
    shift2 := vector3.InitVector3(-5.0, -1.0, -20.0)
    oppositeCorner2 := vector3.InitVector3(2.0, 4.0, -15.5)
    var seed2 int64 = 250
    perlinNoise2 := noise.InitPerlinNoise(1.0, 2.0, 1.0, 0.5, 5, seed2)
    voxelGrid2 := voxel_grid.InitVoxelGrid(0.1, shift2, oppositeCorner2, 0.05, perlinNoise2)

    voxelGrids := []voxel_grid.VoxelGrid{voxelGrid, voxelGrid2}


    // Spheres
    sphere := sphere.InitSphere(vector3.InitVector3(0, 0, -2), 1.0)

    // Lights
    light := light.InitLight(vector3.InitVector3(0.0, 6.0, 0.0), vector3.InitVector3(0.6, 0.6, 0.6))


    // Scene
    s := scene.InitScene(voxelGrids, sphere, camera, light);

    // Render
    image := s.Render(imgSizeY, imgSizeX, 3)

    // Save
    image.SavePPM("tmp.ppm")
}
