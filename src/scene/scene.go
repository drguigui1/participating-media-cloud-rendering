package scene

import (
    //"volumetric-cloud/vector3"
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/ray"
    "volumetric-cloud/vector3"

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
    ray := s.Camera.CreateRay(j, i)

    // Check intersect with Voxel Grid
    _, hasHit, color := s.VoxelGrid.Hit(ray)
    //_, _, hasHit := s.Sphere.Hit(ray)

    // raymarch TODO

    // set pixel
    if hasHit {
        //image.SetPixel(i, j, 255, 111, 0)
        image.SetPixel(i, j, byte(color.X), byte(color.Y), byte(color.Z))
    } else {
        // gradient case
        color := s.RenderGradient(ray)
        image.SetPixel(i, j, byte(color.X * 255.0), byte(color.Y * 255), byte(color.Z * 255))
    }
}

func (s Scene) renderPixel(image img.Img, i, j int, wg *sync.WaitGroup) {

    // create the ray
//    ray := s.Camera.CreateRay(i, j)

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

func (s Scene) RenderGradient(ray ray.Ray) vector3.Vector3 {
    dir := vector3.UnitVector(ray.Direction);
    tmp := 0.5 * (dir.Y + 1.0);
    tmp2 := 1.0 - tmp
    return vector3.AddVector3(vector3.InitVector3(1.0 * tmp2, 1.0 * tmp2, 1.0 * tmp2), vector3.InitVector3(0.35 * tmp, 0.76 * tmp, 0.75 * tmp))
}
