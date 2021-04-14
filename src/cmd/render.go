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
    "volumetric-cloud/random_clouds"
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

        origin := vector3.InitVector3(-5.0, 0.0, 0.0)
        camera := camera.InitCamera(
           aspectRatio,
           fieldOfView,
           imgSizeX,
           imgSizeY,
           origin,
           math.Pi / 4,
           math.Pi / 8,
           0.0,
        )
            // Voxel Grid 2
           // Image 'perlin-worley-2.png'
        shift2 := vector3.InitVector3(-50, 35.0, -60.0)
        oppositeCorner2 := vector3.InitVector3(-25.0, 40.0, -30.0)
        var seed2 int64 = 21
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(1.0, 2.0, 1.0, 2.0, 3, seed2)
        worleyWeight := 0.0
        perlinWeight := 0.5
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift2, oppositeCorner2, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.3, 0.3, 1.0)






        /*
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


         */

        /*
        shift := vector3.InitVector3(-20.0, 25.0, -90.0)
        oppositeCorner := vector3.InitVector3(20.0, 30.0, -50.0)
        var seed2 int64 = 2
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed2)
        worleyWeight := 0.2
        perlinWeight := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.3, 0.6, 1.5)
    */
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


        /*
        fmt.Println("RENDER")
        image := s.Render(imgSizeY, imgSizeX)
         */


        cloudCenter := vector3.InitVector3(-37.5, 37.5, -45.0)
        image := animations.LookCenter(cloudCenter, s, imgSizeX, imgSizeY)
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

        origin := vector3.InitVector3(-5.0, 0.0, 30)
        camera := camera.InitCamera(
           aspectRatio,
           fieldOfView,
           imgSizeX,
           imgSizeY,
           origin,
           math.Pi / 3.0,
           0.0,
           0.0,
        )

        // Voxel Grid 2
        // Image 'perlin-worley-2.png'
        shift := vector3.InitVector3(-30, 35.0, 10.0)
        oppositeCorner := vector3.InitVector3(-5.0, 40.0, 50.0)
        var seed int64 = 21
        worleyNoise := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed)
        perlinNoise := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed)
        worleyWeight := 0.5
        perlinWeight := 0.5
        voxelGrid1 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise, worleyNoise, perlinWeight, worleyWeight, 0.6, 0.6, 1.5)

        /*
        shifté := vector3.InitVector3(-30, 35.0, 10.0)
        oppositeCorner2 := vector3.InitVector3(-5.0, 40.0, 50.0)
        var seed2 int64 = 21
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed)
        worleyWeight2 := 0.5
        perlinWeight2 := 0.5
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise, worleyNoise, perlinWeight, worleyWeight, 0.6, 0.6, 1.5)


         */

/*
        shift2 := vector3.InitVector3(-20.0, 25.0, -90.0)
        oppositeCorner2 := vector3.InitVector3(20.0, 30.0, -50.0)
        var seed2 int64 = 2
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed2)
        worleyWeight2 := 0.2
        perlinWeight2 := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift2, oppositeCorner2, 0.13, perlinNoise2, worleyNoise2, perlinWeight2, worleyWeight2, 0.3, 0.6, 1.5)

        shift3 := vector3.InitVector3(20.0, 50.0, 20.0)
        oppositeCorner3 := vector3.InitVector3(40.0, 80.0, 60)
        var seed3 int64 = 1964
        worleyNoise3 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed3)
        perlinNoise3 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed3)
        worleyWeight3 := 0.2
        perlinWeight3 := 0.6
        voxelGrid3 := voxel_grid.InitVoxelGrid(0.5, shift3, oppositeCorner3, 0.13, perlinNoise3, worleyNoise3, perlinWeight3, worleyWeight3, 0.5, 0.5, 1.5)

*/
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
        voxelGrids := []voxel_grid.VoxelGrid{voxelGrid1}

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
        direct := vector3.InitVector3(0.0, 1.0, 0.0)
        r := ray.InitRay(s.Camera.Origin, direct)

        // animTranslate(ray, picNumber, imgX, imY, nbRpp, step, scene, cam)
        animations.AnimTranslate(r, 10, imgSizeX, imgSizeY, 1.0,  s, camera)
    },
}

var randomRenderCmd = &cobra.Command{
    Use: "randomrender",
    Short: "Generate random clouds and render",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 1200
        imgSizeY := 1000

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(0, 0, 0)
        camera := camera.InitCamera(
            aspectRatio,
            fieldOfView,
            imgSizeX,
            imgSizeY,
            origin,
            math.Pi / 10.0,
            0.0,
            0.0,
        )

        // Init random voxelGrids

        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, 200.0), vector3.InitVector3(0.8, 0.8, 0.8))
        //light := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1}

        // Atmosphere
        ground := sphere.InitSphere(vector3.InitVector3(0.0, -150.0, 0.0), 150.0)
        groundColor := vector3.InitVector3(91.0 / 255.0, 113 / 255.0, 182.0 / 255.0)
        albedo := 0.9
        atmosphere := atmosphere.Atmosphere{
            Ground: ground,
            GroundColor: groundColor,
            GroundAlbedo: albedo,
        }


        // noise
        var defaultSeed int64 = 1
        //freq, freqFactor, amplitude, amplitudeFactor float64, octaves int, seed
        worley := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, defaultSeed)
        perlin := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, defaultSeed)
        worleyW := 0.5
        perlinW := 0.5

        //cloud parameters
        voxelSize := 0.5
        step := 0.2
        sharpness := 0.6
        cloudCoverValue := 0.3
        densityFactor := 1.5
        nbClouds := 20

        //range pos
        posX := []int {-100, 100}
        posY := []int {-100, 100}
        posZ := []int {-100, 100}

        fmt.Println("RAND")

        //freq, freqFactor, amplitude, amplitudeFactor float64, octaves int
        voxelGrids := random_clouds.GenerateRandomClouds(posX, posY, posZ, perlin, worley,
            perlinW, worleyW, voxelSize, step, sharpness, cloudCoverValue, densityFactor,
            nbClouds)

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, atmosphere, 0.4)

        fmt.Println("ANIM")

        // Render

        direct := vector3.InitVector3(0.0, 0.0, 1.0)
        r := ray.InitRay(s.Camera.Origin, direct)
        animations.AnimTranslate(r, 150, imgSizeX, imgSizeY, 1,  s, camera)

        //center := vector3.InitVector3(0.0, 0.0, 0.0)
        //animations.AnimRotation(center, 30.0, imgSizeX, imgSizeY, 100, s)

        //animations.animTranslate(ray, picNumber, imgX, imY, nbRpp, step, scene, cam)
    },
}



func init() {
    cmd.AddCommand(fullRenderCmd)
    cmd.AddCommand(animateRenderCmd)
    cmd.AddCommand(randomRenderCmd)
}
