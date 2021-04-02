package scene

import (
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/background"
    "volumetric-cloud/vector3"

    "math/rand"
    "sync"
)

type Scene struct {
    // TODO change to many VoxelGrid
    VoxelGrids []voxel_grid.VoxelGrid
    Sphere sphere.Sphere
    Camera camera.Camera
    Light light.Light
}

func InitScene(voxelGrids []voxel_grid.VoxelGrid,
               sphere sphere.Sphere,
               camera camera.Camera,
               light light.Light) Scene {
    // compute light transmittance in the voxel grid
    for idx, _ := range voxelGrids {
        voxelGrids[idx].ComputeInsideLightTransparency(light)
    }

    return Scene{
        VoxelGrids: voxelGrids,
        Sphere: sphere,
        Camera: camera,
        Light: light,
    }
}

func (s Scene) Render(imgSizeY, imgSizeX, nbRaysPerPixel int) img.Img {
    image := img.InitImg(imgSizeY, imgSizeX)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY * imgSizeX)

    for i := 0; i < imgSizeY; i += 1 {
        for j := 0; j < imgSizeX; j += 1 {
            //go s.renderPixel(image, i, j, &wg)
            s.renderPixelNoGoroutine(image, i, j, nbRaysPerPixel)
        }
    }

    return image
}

func (s Scene) renderPixelNoGoroutine(image img.Img, i, j, nbRaysPerPixel int) {
    color := vector3.InitVector3(0, 0, 0)
    for k := 0; k < nbRaysPerPixel; k += 1 {
        // create the ray
        // need first column index (j) and then row index (i)
        r := s.Camera.CreateRay(float64(j) + rand.Float64(), float64(i) + rand.Float64())

        // Check intersect with Voxel Grids
        var c vector3.Vector3
        var hasHit bool
        for _, vGrid := range s.VoxelGrids {
            c, hasHit = vGrid.ComputePixelColor(r, s.Light.Color)
            if hasHit {
                break
            }
        }

        // set pixel
        if hasHit {
            // compute pizel color
            color.AddVector3(vector3.InitVector3(c.X, c.Y, c.Z))
        } else {
            // gradient case
            colorG := background.RenderGradient(r)
            color.AddVector3(colorG)
        }
    }

    // divide color vector by nbRaysPerPixel
    color.Div(float64(nbRaysPerPixel))

    // Set the pixel color
    image.SetPixel(i, j, byte(color.X * 255.0), byte(color.Y * 255.0), byte(color.Z * 255.0))
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
