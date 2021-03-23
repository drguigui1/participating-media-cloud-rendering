package scene

import (
    //"volumetric-cloud/vector3"
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"

    "sync"
)

type Scene struct {
    // TODO change to many VoxelGrid
    VoxelGrid voxel_grid.VoxelGrid
    Camera camera.Camera
}

func InitScene(voxelGrid voxel_grid.VoxelGrid, camera camera.Camera) Scene {
    return Scene{
        VoxelGrid: voxelGrid,
        Camera: camera,
    }
}

func (s Scene) Render(imgSizeY, imgSizeX int) img.Img {
    image := img.InitImg(imgSizeY, imgSizeX)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY * imgSizeX)

    for i := 0; i < imgSizeY; i += 1 {
        for j := 0; j < imgSizeX; j += 1 {
            go s.renderPixel(image, i, j, &wg)
            //image.SetPixel(i, j, 200, 100, 20)
        }
    }

    return image
}

func (s Scene) renderPixel(image img.Img, i, j int, wg *sync.WaitGroup) {

    // create the ray
    ray := s.Camera.CreateRay(i, j)

    // Check intersect with Voxel Grid
    _, hasHit, color := s.VoxelGrid.IntersectFaces(ray)

    // raymarch TODO

    // set pixel
    if hasHit {
        image.SetPixel(i, j, byte(color.X), byte(color.Y), byte(color.Z))
    } else {
        image.SetPixel(i, j, 255, 255, 255)
    }

    wg.Done()
}
