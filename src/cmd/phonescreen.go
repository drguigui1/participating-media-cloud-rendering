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
    "volumetric-cloud/atmosphere"
    "volumetric-cloud/sphere"
)

var phoneSunset = &cobra.Command{
    Use: "phonesunset",
    Short: "Generate Phone dimension sunset",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 720
        imgSizeY := 1280

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(0.0, 0.0, 5.0)
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

        // voxelgrid1
        shift := vector3.InitVector3(-20.0, 35.0, -90.0)
        oppositeCorner := vector3.InitVector3(20.0, 40.0, -50.0)
        var seed1 int64 = 2
        worleyNoise1 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed1)
        perlinNoise1 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed1)
        worleyWeight := 0.2
        perlinWeight := 0.6
        voxelGrid1 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise1, worleyNoise1, perlinWeight, worleyWeight, 0.5, 0.6, 1.0)

        //voxelgrid2
        shift2 := vector3.InitVector3(-90.0, 45.0, -200.0)
        oppositeCorner2 := vector3.InitVector3(-60.0, 50.0, -120.0)
        var seed2 int64 = 4
        worleyNoise2 := noise.InitWorleyNoise(0.4, 1.5, 0.7, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, seed2)
        worleyWeight2 := 0.2
        perlinWeight2 := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift2, oppositeCorner2, 0.13, perlinNoise2, worleyNoise2, perlinWeight2, worleyWeight2, 0.4, 0.3, 1.0)

        //voxelgrid3
        shift3 := vector3.InitVector3(100, 50.0, -150.0)
        oppositeCorner3 := vector3.InitVector3(150, 55.0, -110.0)
        var seed3 int64 = 6
        worleyNoise3 := noise.InitWorleyNoise(0.4, 1.5, 0.7, 0.5, 3, seed3)
        perlinNoise3 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, seed3)
        worleyWeight3 := 0.8
        perlinWeight3 := 0.2
        voxelGrid3 := voxel_grid.InitVoxelGrid(0.5, shift3, oppositeCorner3, 0.13, perlinNoise3, worleyNoise3, perlinWeight3, worleyWeight3, 0.4, 0.3, 1.0)

        //voxelgrid4
        shift4 := vector3.InitVector3(80, 80.0, -150.0)
        oppositeCorner4 := vector3.InitVector3(130, 100.0, -80.0)
        var seed4 int64 = 8
        worleyNoise4 := noise.InitWorleyNoise(0.4, 1.5, 0.7, 0.5, 3, seed4)
        perlinNoise4 := noise.InitPerlinNoise(0.1, 2.0, 1.0, 0.5, 3, seed4)
        worleyWeight4 := 0.2
        perlinWeight4 := 0.6
        voxelGrid4 := voxel_grid.InitVoxelGrid(0.5, shift4, oppositeCorner4, 0.13, perlinNoise4, worleyNoise4, perlinWeight4, worleyWeight4, 0.4, 0.5, 2.5)

        //voxelgrid5
        shift5 := vector3.InitVector3(-150, 70.0, -180.0)
        oppositeCorner5 := vector3.InitVector3(-100, 80.0, -110.0)
        var seed5 int64 = 10
        worleyNoise5 := noise.InitWorleyNoise(0.5, 1.5, 0.7, 0.5, 3, seed5)
        perlinNoise5 := noise.InitPerlinNoise(0.1, 2.0, 1.0, 0.5, 3, seed5)
        worleyWeight5 := 0.3
        perlinWeight5 := 0.5
        voxelGrid5 := voxel_grid.InitVoxelGrid(0.5, shift5, oppositeCorner5, 0.13, perlinNoise5, worleyNoise5, perlinWeight5, worleyWeight5, 0.4, 0.4, 1.8)

        //voxelgrid7
        shift7 := vector3.InitVector3(100, 50.0, -250.0)
        oppositeCorner7 := vector3.InitVector3(150, 60.0, -190.0)
        var seed7 int64 = 33
        worleyNoise7 := noise.InitWorleyNoise(0.5, 1.5, 0.7, 0.5, 3, seed7)
        perlinNoise7 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, seed7)
        worleyWeight7 := 0.3
        perlinWeight7 := 0.7
        voxelGrid7 := voxel_grid.InitVoxelGrid(0.5, shift7, oppositeCorner7, 0.13, perlinNoise7, worleyNoise7, perlinWeight7, worleyWeight7, 0.4, 0.4, 2.2)

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
        voxelGrids := []voxel_grid.VoxelGrid{voxelGrid1, voxelGrid2, voxelGrid3, voxelGrid4, voxelGrid5, voxelGrid7}
        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(10.0, 10.0, 10.0), vector3.InitVector3(0.7, 0.7, 0.7))
        light2 := light.InitLight(vector3.InitVector3(0.0, 600000.0, -600000.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1, light2}

        sun := light.InitLight(vector3.InitVector3(0.0, 30, -1000.0), vector3.InitVector3(20.0, 10.0, 10.0))

        // Atmosphere
        ground := sphere.InitSphere(vector3.InitVector3(0.0, -6360005, 0.0), 6360000)

        groundColor := vector3.InitVector3(32 / 255.0, 117 / 255.0, 133 / 255.0)
        albedo := 0.9

        scaleHeightR := 8000.0
        scaleHeightM := 1200.0

        sunImpact := vector3.InitVector3(241.0 / 255.0, 161.0 / 255.0, 109.0 / 255.0)

        betaRayleigh := vector3.InitVector3(0.0000058, 0.0000135, 0.0000331)
        vMie := 0.0000010
        betaMie := vector3.InitVector3(vMie, vMie, vMie)
        atmosphere := atmosphere.InitAtmosphere(
            ground,
            groundColor,
            sunImpact,
            betaRayleigh,
            betaMie,
            albedo,
            6380000,
            sun,
            8,
            8,
            scaleHeightR,
            scaleHeightM,
        )

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, atmosphere, 0.3)

        fmt.Println("RENDER")
        image := s.Render(imgSizeY, imgSizeX)

        fmt.Println("SAVE")
        image.SavePNG("tmp.png")


    },
}

func init() {
    cmd.AddCommand(phoneSunset)
}
