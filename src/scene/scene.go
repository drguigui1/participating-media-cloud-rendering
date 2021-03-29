package scene

import (
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/background"

    "sync"
)

type Scene struct {
    // TODO change to many VoxelGrid
    VoxelGrid voxel_grid.VoxelGrid
    Sphere sphere.Sphere
    Camera camera.Camera
    Light light.Light
}

func InitScene(voxelGrid voxel_grid.VoxelGrid,
               sphere sphere.Sphere,
               camera camera.Camera,
               light light.Light) Scene {
    // compute light transmittance in the voxel grid
    voxelGrid.ComputeInsideLightTransparency(light)
    return Scene{
        VoxelGrid: voxelGrid,
        Sphere: sphere,
        Camera: camera,
        Light: light,
    }
}

func (s Scene) Render(imgSizeY, imgSizeX int) img.Img {
    image := img.InitImg(imgSizeY, imgSizeX)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY * imgSizeX)

    for i := 0; i < imgSizeY; i += 1 {
        for j := 0; j < imgSizeX; j += 1 {
            //go s.renderPixel(image, i, j, &wg)
            s.renderPixelNoGoroutine(image, i, j)
            //image.SetPixel(i, j, 200, 100, 20)
        }
    }

    return image
}

func (s Scene) renderPixelNoGoroutine(image img.Img, i, j int) {
    // create the ray
    // need first column index (j) and then row index (i)
    r := s.Camera.CreateRay(j, i)

    // Check intersect with Voxel Grid
    color, hasHit := s.VoxelGrid.RenderPixel(r, s.Light.Color)

    // set pixel
    if hasHit {
        // compute pizel color
        image.SetPixel(i, j, byte(color.X * 255.0), byte(color.Y * 255.0), byte(color.Z * 255.0))
    } else {
        // gradient case
        colorG := background.RenderGradient(r)
        image.SetPixel(i, j, byte(colorG.X * 255.0), byte(colorG.Y * 255), byte(colorG.Z * 255))
    }
}

func (s Scene) renderPixel(image img.Img, i, j int, wg *sync.WaitGroup) {
    // create the ray
    // ray := s.Camera.CreateRay(i, j)

    // Check intersect with Voxel Grid
    //_, hasHit, color := s.VoxelGrid.IntersectFaces(ray, i, j)

    // raymarch TODO

    // set pixel
//    if hasHit {
//        image.SetPixel(i, j, byte(color.X), byte(color.Y), byte(color.Z))
//    } else {
//        image.SetPixel(i, j, 255, 255, 255)
//    }
//
//    wg.Done()
}
