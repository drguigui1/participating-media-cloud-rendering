package main

import (
    "math"
    "fmt"

    "volumetric-cloud/camera"
    "volumetric-cloud/light"
    "volumetric-cloud/noise"
    "volumetric-cloud/scene"
    "volumetric-cloud/vector3"
    "volumetric-cloud/voxel_grid"
)

func main() {
    imgSizeX := 1200
    imgSizeY := 1000

    // Camera
    aspectRatio := float64(imgSizeX) / float64(imgSizeY)
    fieldOfView := math.Pi / 2
/*    origin := vector3.InitVector3(-12, 6, -13)
    camera := camera.InitCamera(
       aspectRatio,
       fieldOfView,
       imgSizeX,
       imgSizeY,
       origin,
       -math.Pi / 8,
       -math.Pi / 2,
       0.0,
    )*/

    origin := vector3.InitVector3(0, 15, 5)
    camera := camera.InitCamera(
       aspectRatio,
       fieldOfView,
       imgSizeX,
       imgSizeY,
       origin,
       math.Pi / 8,
       0.0,
       0.0,
    )

    // Voxel Grid 1
    shift := vector3.InitVector3(-20.0, 35.0, -90.0)
    oppositeCorner := vector3.InitVector3(20.0, 40.0, -60.0)
    var seed int64 = 42
    perlinNoise := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, seed)
    voxelGrid := voxel_grid.InitVoxelGrid(0.2, shift, oppositeCorner, 0.12, perlinNoise)

    // Voxel Grid 2
    shift2 := vector3.InitVector3(-45.0, 35.0, -120.0)
    oppositeCorner2 := vector3.InitVector3(-5.0, 40.0, -100.0)
    var seed2 int64 = 21
    perlinNoise2 := noise.InitPerlinNoise(0.3, 2.0, 1.0, 0.5, 4, seed2)
    voxelGrid2 := voxel_grid.InitVoxelGrid(0.2, shift2, oppositeCorner2, 0.13, perlinNoise2)

    // Voxel Grid 3
    shift3 := vector3.InitVector3(20.0, 35.0, -130.0)
    oppositeCorner3 := vector3.InitVector3(40.0, 40.0, -110.0)
    var seed3 int64 = 200
    perlinNoise3 := noise.InitPerlinNoise(0.3, 2.0, 1.0, 0.5, 4, seed3)
    voxelGrid3 := voxel_grid.InitVoxelGrid(0.2, shift3, oppositeCorner3, 0.13, perlinNoise3)

    // Voxel Grid 4
    shift4 := vector3.InitVector3(-80.0, 35.0, -80.0)
    oppositeCorner4 := vector3.InitVector3(-50.0, 40.0, -70.0)
    var seed4 int64 = 10
    perlinNoise4 := noise.InitPerlinNoise(0.3, 2.0, 1.0, 0.4, 3, seed4)
    voxelGrid4 := voxel_grid.InitVoxelGrid(0.2, shift4, oppositeCorner4, 0.13, perlinNoise4)

    // IMPORTANT
    //
    // First condition:
    // (oppositeCorner.X - shift.X) / voxelSize  => MUST BE AN INTEGER (NO FLOAT)
    // (oppositeCorner.Y - shift.Y) / voxelSize  => MUST BE AN INTEGER (NO FLOAT)
    // (oppositeCorner.Z - shift.Z) / voxelSize  => MUST BE AN INTEGER (NO FLOAT)
    //
    // Second condition:
    // shift.X < oppositeCorner.X &&
    // shift.Y < oppositeCorner.Y &&
    // shift.Z < oppositeCorner.Z
    fmt.Println("VOXEL")
    voxelGrids := []voxel_grid.VoxelGrid{voxelGrid, voxelGrid2, voxelGrid3, voxelGrid4}

    // Lights
    light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, 0.0), vector3.InitVector3(0.25, 0.25, 0.25))
    light2 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.2, 0.2, 0.2))
    //light3 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.3, 0.3, 0.3))
    //light4 := light.InitLight(vector3.InitVector3(100.0, 100.0, 100.0), vector3.InitVector3(0.4, 0.4, 0.4))

    lights := []light.Light{light1, light2}

    // Scene
    fmt.Println("SCENE")
    s := scene.InitScene(voxelGrids, camera, lights)

    fmt.Println("RENDER")
    // Render
    image := s.Render(imgSizeY, imgSizeX, 1)

    fmt.Println("SAVE")
    // Save
    image.SavePPM("tmp.ppm")




}
