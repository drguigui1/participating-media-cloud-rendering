package main

//import (
//    "math"
//
//    "volumetric-cloud/voxel_grid"
//    "volumetric-cloud/scene"
//    "volumetric-cloud/sphere"
//    "volumetric-cloud/camera"
//    "volumetric-cloud/vector3"
//)
//
func main() {
//    imgSizeX := 600
//    imgSizeY := 500
//
//    // Camera
//    aspectRatio := float64(imgSizeX) / float64(imgSizeY)
//    fieldOfView := math.Pi / 2.0
//    origin := vector3.InitVector3(0.0, 2.0, 0.0)
//    camera := camera.InitCamera(
//        aspectRatio,
//        fieldOfView,
//        imgSizeX,
//        imgSizeY,
//        origin,
//        -math.Pi / 4.0,
//        0.0,
//        0.0,
//    )
//
//    // Voxel Grid
//    shift := vector3.InitVector3(0.0, 0.0, -4.0)
//    voxelGrid := voxel_grid.InitVoxelGrid(0.2, 10, 10, 10, shift)
//
//    // Spheres
//    sphere := sphere.InitSphere(vector3.InitVector3(0, 0, -2), 1.0)
//
//    // Scene
//    s := scene.InitScene(voxelGrid, sphere, camera);
//
//    // Render
//    image := s.Render(imgSizeY, imgSizeX)
//
//    // Save
//    image.SavePPM("tmp.ppm")
//
}
