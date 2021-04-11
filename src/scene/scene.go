package scene

import (
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"
    "volumetric-cloud/light"
    "volumetric-cloud/background"
    "volumetric-cloud/vector3"

    "math/rand"
    "sync"
    "fmt"
)

type Pixel struct {
    I int
    J int
    Intensity []float64
    BackgroundColorImpact vector3.Vector3
}

type Scene struct {
    // TODO change to many VoxelGrid
    VoxelGrids []voxel_grid.VoxelGrid
    Camera camera.Camera
    Lights []light.Light

    RainyNess float64 // close to 0 will make a whiter cloud
    MinColor []float64 // rgb (size 3/4)
    MaxColor []float64 // rgb
    Pixels []Pixel
}

func InitScene(voxelGrids []voxel_grid.VoxelGrid,
               camera camera.Camera,
               lights []light.Light,
               rainyNess float64) Scene {
    // compute light transmittance in the voxel grid
    for idx, _ := range voxelGrids {
        voxelGrids[idx].ComputeInsideLightTransparency(lights)
    }

    return Scene{
        VoxelGrids: voxelGrids,
        Camera: camera,
        Lights: lights,
        RainyNess: rainyNess,
        MinColor: []float64{0.0, 0.0, 0.0},
        MaxColor: []float64{0.0, 0.0, 0.0},
        Pixels: make([]Pixel, 0),
    }
}

func (s *Scene) Render(imgSizeY, imgSizeX, nbRaysPerPixel int) img.Img {
    image := img.InitImg(imgSizeX, imgSizeY)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY)

    for i := 0; i < imgSizeY; i += 1 {
        if i == 620 {
            //fmt.Println("BREAK")
        }
        s.renderImageSizeY(image, i, imgSizeX, nbRaysPerPixel, nil)
    }

    // i == 620
    // j == 300

    //for j := 300; j < 700; j += 1 {
    //    image.SetPixel(j, 620, 255, 0, 0, 255)
    //}
    //wg.Wait()

    // Remap cloud values
    fmt.Println(s.MinColor)
    fmt.Println(s.MaxColor)
    for _, p := range s.Pixels {
        colorX := (p.Intensity[0] - s.MinColor[0]) / (s.MaxColor[0] - s.MinColor[0])
        colorY := (p.Intensity[1] - s.MinColor[1]) / (s.MaxColor[1] - s.MinColor[1])
        colorZ := (p.Intensity[2] - s.MinColor[2]) / (s.MaxColor[2] - s.MinColor[2])

        color := vector3.InitVector3(colorX, colorY, colorZ)
        //p.BackgroundColorImpact.Sub(0.7)
        color.AddVector3(p.BackgroundColorImpact)
        color.Div(float64(nbRaysPerPixel))
        color.Clamp(0.0, 1.0)
        image.SetPixel(p.J, p.I, uint8(color.X * 255.0), uint8(color.Y * 255.0), uint8(color.Z * 255.0), uint8(255))
    }

    return image
}

func (s *Scene) renderImageSizeY(image img.Img, i, imgSizeX, nbRaysPerPixel int, wg *sync.WaitGroup) {
    for j := 0; j < imgSizeX; j += 1 {
        color := vector3.InitVector3(0, 0, 0)
        if j == 300 {
            //fmt.Println("BREAK")
        }

        var hasOneRayHit bool = false
        accColor := vector3.InitVector3(0, 0, 0)
        backgroundColorImpact := vector3.InitVector3(0, 0, 0)

        for k := 0; k < nbRaysPerPixel; k += 1 {
            // create the ray
            r := s.Camera.CreateRay(float64(j) + rand.Float64(), float64(i) + rand.Float64())

            var accC vector3.Vector3

            var accTransparency float64 = 1.0
            var accT float64

            var hasHit bool
            hasOneHit := false

            for _, vGrid := range s.VoxelGrids {
                accC, accT, hasHit = vGrid.ComputePixelColor(r, s.Lights[0].Color, s.RainyNess)
                if !hasHit {
                    continue
                }

                hasOneHit = true
                hasOneRayHit = true

                // accumulate transparency
                accTransparency *= accT

                // accumulate color
                accColor.AddVector3(accC)
            }

            // get background impact
            backgroundColor := background.RenderGradient(r)

            // set pixel
            if hasOneHit {
                //accColor.Mul(s.RainyNess)
                accColor.Mul(s.RainyNess)
                // compute pizel color
                backgroundColorImpact.AddVector3(vector3.MulVector3Scalar(backgroundColor, accTransparency))
                //accColor.AddVector3(backgroundColorImpact)
                // accColor.Clamp(0.0, 1.0)
                //color.AddVector3(vector3.InitVector3(accColor.X, accColor.Y, accColor.Z))
                //color = vector3.InitVector3(0.0, 1.0, 0.0)
            } else {
                // gradient case
                color.AddVector3(backgroundColor)
            }
        }

        // divide color vector by nbRaysPerPixel
        color.Div(float64(nbRaysPerPixel))

        if hasOneRayHit {
            // Build min / max / tuple slice (i, j, Intensity)
            //if hasOneHit {
            if accColor.X < s.MinColor[0] {
                s.MinColor[0] = accColor.X
            }
            if accColor.Y < s.MinColor[1] {
                s.MinColor[1] = accColor.Y
            }
            if accColor.Z < s.MinColor[2] {
                s.MinColor[2] = accColor.Z
            }

            if accColor.X > s.MaxColor[0] {
                s.MaxColor[0] = accColor.X
            }
            if accColor.Y > s.MaxColor[1] {
                s.MaxColor[1] = accColor.Y
            }
            if accColor.Z > s.MaxColor[2] {
                s.MaxColor[2] = accColor.Z
            }

            s.Pixels = append(s.Pixels, Pixel{
                I: i,
                J: j,
                Intensity: []float64{accColor.X, accColor.Y, accColor.Z},
                BackgroundColorImpact: backgroundColorImpact,
            })
        } else {
            image.SetPixel(j, i, uint8(color.X * 255.0), uint8(color.Y * 255.0), uint8(color.Z * 255.0), uint8(255))
        }
    }

    if wg != nil {
        wg.Done()
    }
}
