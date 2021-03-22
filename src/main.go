package main

import (
    //"fmt"

    //"volumetric-cloud/vector3"
    //"volumetric-cloud/ray"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/scene"
    "volumetric-cloud/camera"
)

func main() {
    imgSizeX := 256
    imgSizeY := 256
    s := scene.InitScene(voxel_grid.VoxelGrid{}, camera.Camera{});
    image := s.Render(imgSizeY, imgSizeX)
    image.SavePPM("tmp.ppm")

}
