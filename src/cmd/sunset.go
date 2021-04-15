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
	"volumetric-cloud/animations"
	//"volumetric-cloud/ray"
)


var sunset1 = &cobra.Command{
	Use: "sunset1",
	Short: "Generate Sunset1",
	Run: func(cmd *cobra.Command, args []string) {
		imgSizeX := 1200
		imgSizeY := 1000

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
			math.Pi / 10,
			0.0,
			0.0,
		)

		// voxelSize float64, shift vector3.Vector3, oppositeCorner vector3.Vector3, step float64, perlinNoise noise.PerlinNoise, worleyNoise noise.WorleyNoise, perlinWeight, worleyWeight, sharpness, cloudCoverVal, densityFactor float64
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

		/*
		//voxelgrid6 marche pas encore
		shift6 := vector3.InitVector3(-130, 130.0, -1000.0)
		oppositeCorner6 := vector3.InitVector3(130, 200.0, -985.0)
		var seed6 int64 = 14
		worleyNoise6 := noise.InitWorleyNoise(0.5, 1.5, 0.7, 0.5, 3, seed6)
		perlinNoise6 := noise.InitPerlinNoise(0.1, 2.0, 1.0, 0.5, 3, seed6)
		worleyWeight6 := 0.3
		perlinWeight6 := 0.5
		voxelGrid6 := voxel_grid.InitVoxelGrid(0.5, shift6, oppositeCorner6, 0.13, perlinNoise6, worleyNoise6, perlinWeight6, worleyWeight6, 0.3, 0.6, 1.0)


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
		voxelGrids := []voxel_grid.VoxelGrid{voxelGrid1, voxelGrid2, voxelGrid3, voxelGrid4, voxelGrid5}
		_ = voxelGrid1
		_ = voxelGrid2
		_ = voxelGrid3
		fmt.Println("VOXEL")

		// Lights
		light1 := light.InitLight(vector3.InitVector3(10.0, 10.0, 10.0), vector3.InitVector3(0.7, 0.7, 0.7))
		light2 := light.InitLight(vector3.InitVector3(0.0, 600000.0, -600000.0), vector3.InitVector3(0.7, 0.7, 0.7))
		lights := []light.Light{light1, light2}

		sun := light.InitLight(vector3.InitVector3(0.0, 30, -1000.0), vector3.InitVector3(20.0, 10.0, 10.0))

		// Atmosphere
		//ground := sphere.InitSphere(vector3.InitVector3(0.0, -6360005, 0.0), 6360000)
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

var sunsetanim = &cobra.Command{
	Use: "sunsetanim",
	Short: "Generate Sunset1 animation",
	Run: func(cmd *cobra.Command, args []string) {
		imgSizeX := 1200
		imgSizeY := 1000

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
			math.Pi / 10,
			0.0,
			0.0,
		)

		// voxelSize float64, shift vector3.Vector3, oppositeCorner vector3.Vector3, step float64, perlinNoise noise.PerlinNoise, worleyNoise noise.WorleyNoise, perlinWeight, worleyWeight, sharpness, cloudCoverVal, densityFactor float64
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
		shift7 := vector3.InitVector3(100, 80.0, 100.0)
		oppositeCorner7 := vector3.InitVector3(150, 100.0, 160.0)
		var seed7 int64 = 16
		worleyNoise7 := noise.InitWorleyNoise(0.4, 1.5, 0.7, 0.5, 3, seed7)
		perlinNoise7 := noise.InitPerlinNoise(0.1, 2.0, 1.0, 0.5, 3, seed7)
		worleyWeight7 := 0.2
		perlinWeight7 := 0.6
		voxelGrid7 := voxel_grid.InitVoxelGrid(0.5, shift7, oppositeCorner7, 0.13, perlinNoise7, worleyNoise7, perlinWeight7, worleyWeight7, 0.4, 0.5, 2.5)

		//voxelgrid8
		shift8 := vector3.InitVector3(-100, 70.0, 90.0)
		oppositeCorner8 := vector3.InitVector3(-50, 80.0, 150.0)
		var seed8 int64 = 20
		worleyNoise8 := noise.InitWorleyNoise(0.4, 1.5, 0.7, 0.5, 3, seed8)
		perlinNoise8 := noise.InitPerlinNoise(0.1, 2.0, 1.0, 0.5, 3, seed8)
		worleyWeight8 := 0.2
		perlinWeight8 := 0.6
		voxelGrid8 := voxel_grid.InitVoxelGrid(0.5, shift8, oppositeCorner8, 0.13, perlinNoise8, worleyNoise8, perlinWeight8, worleyWeight8, 0.4, 0.4, 1.8)


		/*
			//voxelgrid6 marche pas encore
			shift6 := vector3.InitVector3(-130, 130.0, -1000.0)
			oppositeCorner6 := vector3.InitVector3(130, 200.0, -985.0)
			var seed6 int64 = 14
			worleyNoise6 := noise.InitWorleyNoise(0.5, 1.5, 0.7, 0.5, 3, seed6)
			perlinNoise6 := noise.InitPerlinNoise(0.1, 2.0, 1.0, 0.5, 3, seed6)
			worleyWeight6 := 0.3
			perlinWeight6 := 0.5
			voxelGrid6 := voxel_grid.InitVoxelGrid(0.5, shift6, oppositeCorner6, 0.13, perlinNoise6, worleyNoise6, perlinWeight6, worleyWeight6, 0.3, 0.6, 1.0)


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
		voxelGrids := []voxel_grid.VoxelGrid{voxelGrid1, voxelGrid2, voxelGrid3, voxelGrid4, voxelGrid5, voxelGrid7, voxelGrid8}
		_ = voxelGrid1
		_ = voxelGrid2
		_ = voxelGrid3
		fmt.Println("VOXEL")

		// Lights
		light1 := light.InitLight(vector3.InitVector3(10.0, 10.0, 10.0), vector3.InitVector3(0.7, 0.7, 0.7))
		light2 := light.InitLight(vector3.InitVector3(0.0, 600000.0, -600000.0), vector3.InitVector3(0.7, 0.7, 0.7))
		lights := []light.Light{light1, light2}

		sun := light.InitLight(vector3.InitVector3(0.0, 30, -1000.0), vector3.InitVector3(20.0, 10.0, 10.0))

		// Atmosphere
		//ground := sphere.InitSphere(vector3.InitVector3(0.0, -6360005, 0.0), 6360000)
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


		/*
		fmt.Println("RENDER")
		image := s.Render(imgSizeY, imgSizeX)

		fmt.Println("SAVE")
		image.SavePNG("tmp.png")
		 */

		//direct := vector3.InitVector3(0.0, 0.0, -1.0)
		//r := ray.InitRay(s.Camera.Origin, direct)
		center := vector3.InitVector3(0.0, 0.0, 0.0)
		// animTranslate(ray, picNumber, imgX, imY, nbRpp, step, scene, cam)
		// animations.AnimTranslate(r, 300, imgSizeX, imgSizeY, 0.5,  &s)
		animations.AnimRotation(center, 80, imgSizeX, imgSizeY, 300, s)

	},
}

func init() {
	cmd.AddCommand(sunset1)
	cmd.AddCommand(sunsetanim)
}
