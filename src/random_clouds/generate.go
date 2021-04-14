package random_clouds

import (
    "math/rand"
    "sort"
  //  "time"

    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/noise"
    "volumetric-cloud/vector3"
)

/*
var (
    minX = -100
    maxX = 100
    minY = 30
    maxY = 40
    minZ = -300
    maxZ = -10
)

func GenerateInRange(min, max int, minRange int) []float64 {
    vals := []float64 { float64(rand.Intn(maxX - minX) + minX), float64(rand.Intn(maxX - minX) + minX) }
    sort.Float64s(vals)

    for (vals[1] - vals[0]) < float64(minRange) {
        vals = []float64 { float64(rand.Intn(maxX - minX) + minX), float64(rand.Intn(maxX - minX) + minX) }
        sort.Float64s(vals)
    }

    return vals
}

func GenerateRandomClouds(nbClouds int, minRange int) []voxel_grid.VoxelGrid {
   voxelGrids := make([]voxel_grid.VoxelGrid, nbClouds)

    for i := 0; i < nbClouds; i++ {
        rand.Seed(time.Now().UnixNano())

        shiftX := GenerateInRange(minX, maxX, minRange)
        shiftY := GenerateInRange(minY, maxY, minRange)
        shiftZ := GenerateInRange(minZ, maxZ, minRange)

        perlinNoise := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, int64(i))
        _ = perlinNoise
        shift := vector3.InitVector3(shiftX[0], shiftY[0], shiftZ[0])
        oppositeCorner := vector3.InitVector3(shiftX[1], shiftY[1], shiftZ[1])

        //voxelGrids[i] = voxel_grid.InitVoxelGrid(0.2, shift, oppositeCorner, 0.15, perlinNoise, 0.8, 0.3, 2.0)
        fmt.Println("shift")
        fmt.Println(shift)

        fmt.Println("opposite")
        fmt.Println(oppositeCorner)
    }

    return voxelGrids
}
*/

func GenerateInRange(min, max int) []float64 {
    vals := []float64 { float64(rand.Intn(max - min) + max), float64(rand.Intn(max - min) + min) }
    sort.Float64s(vals)
    return vals
}

func GenerateRandomClouds(posX, posY, posZ []int, p noise.PerlinNoise,
        w noise.WorleyNoise, pW, wW, voxelSize, step, sharpness, cCV,
        densityFactor float64, nbClouds int) []voxel_grid.VoxelGrid {
    voxelGrids := make([]voxel_grid.VoxelGrid, nbClouds)

    for i := 0; i < nbClouds; i++ {
        rand.Seed(50)

        shiftX := GenerateInRange(posX[0], posX[1])
        shiftY := GenerateInRange(posY[0], posY[1])
        shiftZ := GenerateInRange(posZ[0], posZ[1])

        p.SetSeed(int64(i))
        w.SetSeed(int64(i))

        shift := vector3.InitVector3(shiftX[0], shiftY[0], shiftZ[0])
        oppositeCorner := vector3.InitVector3(shiftX[1], shiftY[1], shiftZ[1])

        voxelGrids[i] = voxel_grid.InitVoxelGrid(voxelSize, shift, oppositeCorner,
            step, p, w, pW, wW, sharpness, cCV, densityFactor)
    }

    return voxelGrids
}
