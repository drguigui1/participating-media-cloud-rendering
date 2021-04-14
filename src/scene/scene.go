package scene

import (
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"
    "volumetric-cloud/light"
//    "volumetric-cloud/background"
    "volumetric-cloud/vector3"
    "volumetric-cloud/atmosphere"
    "volumetric-cloud/ray"

    "sync"
    "fmt"
    "math"
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
    Atmosphere atmosphere.Atmosphere

    RainyNess float64 // close to 0 will make a whiter cloud
    MinColor []float64 // rgb (size 3/4)
    MaxColor []float64 // rgb
    Pixels []Pixel
}

func InitScene(voxelGrids []voxel_grid.VoxelGrid,
               camera camera.Camera,
               lights []light.Light,
               atmosphere atmosphere.Atmosphere,
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
        Atmosphere: atmosphere,
        MinColor: []float64{0.0, 0.0, 0.0},
        MaxColor: []float64{0.0, 0.0, 0.0},
        Pixels: make([]Pixel, 0),
    }
}

func (s *Scene) Render(imgSizeY, imgSizeX int) img.Img {
    image := img.InitImg(imgSizeX, imgSizeY)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY)

    for i := 0; i < imgSizeY; i += 1 {
        s.renderImageSizeY(image, i, imgSizeX, nil)
    }

    //wg.Wait()

    // Remap cloud values
    fmt.Println(s.MinColor)
    fmt.Println(s.MaxColor)
    for _, p := range s.Pixels {
        colorX := (p.Intensity[0] - s.MinColor[0]) / (s.MaxColor[0] - s.MinColor[0])
        colorY := (p.Intensity[1] - s.MinColor[1]) / (s.MaxColor[1] - s.MinColor[1])
        colorZ := (p.Intensity[2] - s.MinColor[2]) / (s.MaxColor[2] - s.MinColor[2])

        color := vector3.InitVector3(colorX, colorY, colorZ)
        color.X *= 241.0 / 255.0
        color.Y *= 161.0 / 255.0
        color.Z *= 109.0 / 255.0
        color.AddVector3(p.BackgroundColorImpact)
        color.Clamp(0.0, 1.0)
        image.SetPixel(p.J, p.I, uint8(color.X * 255.0), uint8(color.Y * 255.0), uint8(color.Z * 255.0), uint8(255))
    }

    return image
}

func (s *Scene) renderImageSizeY(image img.Img, i, imgSizeX int, wg *sync.WaitGroup) {
    for j := 0; j < imgSizeX; j += 1 {
        // create the ray
        r := s.Camera.CreateRay(float64(j), float64(i))

        t, _, hasHit := s.Atmosphere.Ground.Hit(r)
        if !hasHit {
            t = -1.0
        }

        // Render
        accColor, backgroundColorImpact, hasOneHit := s.render(r, t)

        if hasOneHit {
            // Build min / max / tuple slice (i, j, Intensity)
            s.MinColor[0] = math.Min(accColor.X, s.MinColor[0])
            s.MinColor[1] = math.Min(accColor.Y, s.MinColor[1])
            s.MinColor[2] = math.Min(accColor.Z, s.MinColor[2])

            s.MaxColor[0] = math.Max(accColor.X, s.MaxColor[0])
            s.MaxColor[1] = math.Max(accColor.Y, s.MaxColor[1])
            s.MaxColor[2] = math.Max(accColor.Z, s.MaxColor[2])

            s.Pixels = append(s.Pixels, Pixel{
                I: i,
                J: j,
                Intensity: []float64{accColor.X, accColor.Y, accColor.Z},
                BackgroundColorImpact: backgroundColorImpact,
            })
        } else {
            image.SetPixel(j, i, uint8(accColor.X * 255.0), uint8(accColor.Y * 255.0), uint8(accColor.Z * 255.0), uint8(255))
        }
    }

    if wg != nil {
        wg.Done()
    }
}

func (s *Scene) render(r ray.Ray, tGround float64) (vector3.Vector3, vector3.Vector3, bool) {
    accColor := vector3.InitVector3(0, 0, 0)
    backgroundColorImpact := vector3.InitVector3(0, 0, 0)

    var accTransparency float64 = 1.0
    hasOneHit := false
    lastPts := vector3.InitVector3(r.Origin.X, r.Origin.Y, r.Origin.Z)

    for _, vGrid := range s.VoxelGrids {
        tVGrid, hasHitVoxel, _ := vGrid.Hit(r)
        if !hasHitVoxel {
            continue
        }

        if tGround > 0 && tVGrid > tGround {
            continue
        }


        var accC vector3.Vector3
        var accT float64
        accC, accT, _, lastPts = vGrid.ComputePixelColor(r, s.Lights[0].Color, s.RainyNess, tGround)

        hasOneHit = true

        // accumulate transparency
        accTransparency *= accT

        // accumulate color
        accColor.AddVector3(accC)
    }

    // get background impact
    //backgroundColor := background.RenderGradient(r)
    var backgroundColor vector3.Vector3

    // Call Rayleigh and Mie to compute the background color
    backgroundColor = s.Atmosphere.ComputeRayleighMie(r)

    if tGround > 0 {
        p := r.RayAt(tGround)
        backgroundColor = s.Atmosphere.Ground.ComputeDiffuseGroundColor(
            s.Lights,
            s.Atmosphere.GroundColor,
            p,
            s.Atmosphere.GroundAlbedo,
        )
    }

    // set pixel
    if hasOneHit {
        accColor.Mul(s.RainyNess)
        backgroundColor = s.Atmosphere.ComputeRayleighMie(ray.InitRay(lastPts, r.Direction))
        //_ = lastPts
        //backgroundColor = s.Atmosphere.ComputeRayleighMie(r)
        backgroundColorImpact.AddVector3(vector3.MulVector3Scalar(backgroundColor, accTransparency))
        return accColor, backgroundColorImpact, hasOneHit
    }

    return backgroundColor, vector3.Vector3{}, hasOneHit
}
