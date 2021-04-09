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
)

var renderCmd = &cobra.Command{
    Use: "render",
    Short: "Different command for rendering clouds",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

var fullRenderCmd = &cobra.Command{
    Use: "fullrender",
    Short: "Generate clouds and render them",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 1200
        imgSizeY := 1000

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(0, 15, 5)
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

        // Voxel Grid 1
        shift := vector3.InitVector3(-20.0, 35.0, -90.0)
        oppositeCorner := vector3.InitVector3(20.0, 40.0, -60.0)
        var seed int64 = 42
        perlinNoise := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, seed)
        voxelGrid := voxel_grid.InitVoxelGrid(0.2, shift, oppositeCorner, 0.12, perlinNoise, 0.6, 0.3, 2.0)

        // Voxel Grid 2
        shift2 := vector3.InitVector3(-45.0, 35.0, -120.0)
        oppositeCorner2 := vector3.InitVector3(-5.0, 40.0, -100.0)
        var seed2 int64 = 21
        perlinNoise2 := noise.InitPerlinNoise(0.3, 2.0, 1.0, 0.5, 4, seed2)
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.2, shift2, oppositeCorner2, 0.13, perlinNoise2, 0.6, 0.3, 1.5)

        voxelGrids := []voxel_grid.VoxelGrid{voxelGrid, voxelGrid2}
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
        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, -200.0), vector3.InitVector3(0.9, 0.9, 0.9))
        light2 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1, light2}

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, 1.0)

        fmt.Println("RENDER")
        // Render
        image := s.Render(imgSizeY, imgSizeX, 1)

        fmt.Println("SAVE")
        // Save
        image.SavePNG("tmp.png")

    },
}

var loadRenderCmd = &cobra.Command{
    Use: "loadrender",
    Short: "Load clouds and render them",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        // load data
        // build the scene
        // launch render
    },
}

func init() {
    renderCmd.AddCommand(fullRenderCmd)
    renderCmd.AddCommand(loadRenderCmd)

    cmd.AddCommand(renderCmd)
}
