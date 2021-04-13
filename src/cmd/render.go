package cmd

import (
    "math"
    "fmt"

    "github.com/spf13/cobra"

    "volumetric-cloud/camera"
    "volumetric-cloud/light"
    "volumetric-cloud/noise"
    "volumetric-cloud/scene"
    "volumetric-cloud/vector3"
    "volumetric-cloud/voxel_grid"
    //"volumetric-cloud/random_clouds"
    "volumetric-cloud/atmosphere"
    "volumetric-cloud/sphere"
    "volumetric-cloud/animations"
    "volumetric-cloud/ray"
)

var fullRenderCmd = &cobra.Command{
    Use: "fullrender",
    Short: "Generate clouds and render them",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 1200
        imgSizeY := 1000

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(0.0, 0.0, 5)
        camera := camera.InitCamera(
           aspectRatio,
           fieldOfView,
           imgSizeX,
           imgSizeY,
           origin,
           0.0,
           0.0,
           0.0,
        )

        shift := vector3.InitVector3(-20.0, 25.0, -90.0)
        oppositeCorner := vector3.InitVector3(20.0, 30.0, -50.0)
        var seed2 int64 = 2
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed2)
        worleyWeight := 0.2
        perlinWeight := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.3, 0.6, 1.5)

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
        voxelGrids := []voxel_grid.VoxelGrid{voxelGrid2}

        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, 200.0), vector3.InitVector3(0.8, 0.8, 0.8))
        light2 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1, light2}

        // Atmosphere
        ground := sphere.InitSphere(vector3.InitVector3(0.0, -150.0, 0.0), 150.0)
        groundColor := vector3.InitVector3(91.0 / 255.0, 113 / 255.0, 182.0 / 255.0)
        albedo := 0.9
        atmosphere := atmosphere.Atmosphere{
            Ground: ground,
            GroundColor: groundColor,
            GroundAlbedo: albedo,
        }

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, atmosphere, 0.3)

        fmt.Println("RENDER")
        image := s.Render(imgSizeY, imgSizeX)

        fmt.Println("SAVE")
        image.SavePNG("tmp.png")
    },
}

var animateRenderCmd = &cobra.Command{
    Use: "animate",
    Short: "Render animation",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 1200
        imgSizeY := 1000

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(-5.0, 0.0, 0.0)
        camera := camera.InitCamera(
           aspectRatio,
           fieldOfView,
           imgSizeX,
           imgSizeY,
           origin,
           0.0,
           0.0,
           0.0,
        )

        // Voxel Grid 1
        // Image 'perlin-worley-3.png'
        shift := vector3.InitVector3(-20.0, 35.0, -50.0)
        oppositeCorner := vector3.InitVector3(20.0, 40.0, -20.0)
        var seed2 int64 = 4
        worleyNoise2 := noise.InitWorleyNoise(0.4, 1.5, 0.7, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, seed2)
        worleyWeight := 0.1
        perlinWeight := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.3, 0.6, 1.5)

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
        voxelGrids := []voxel_grid.VoxelGrid{voxelGrid2}

        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, 200.0), vector3.InitVector3(0.8, 0.8, 0.8))
        light2 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1, light2}

        // Atmosphere
        ground := sphere.InitSphere(vector3.InitVector3(0.0, -150.0, 0.0), 150.0)
        groundColor := vector3.InitVector3(91.0 / 255.0, 113 / 255.0, 182.0 / 255.0)
        albedo := 0.9
        atmosphere := atmosphere.Atmosphere{
            Ground: ground,
            GroundColor: groundColor,
            GroundAlbedo: albedo,
        }

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, atmosphere, 0.4)

        fmt.Println("ANIM")

        // Render
        direct := vector3.InitVector3(0.0, 0.0, 1.0)
        r := ray.InitRay(s.Camera.Origin, direct)

        // animTranslate(ray, picNumber, imgX, imY, nbRpp, step, scene, cam)
        animations.AnimTranslate(r, 100, imgSizeX, imgSizeY, 0.5,  s, camera)
    },
}

func init() {
    cmd.AddCommand(fullRenderCmd)
    cmd.AddCommand(animateRenderCmd)
}
