package random_clouds

import (
    "math/rand"
    "sort"
    "fmt"
    "time"

    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/noise"
    "volumetric-cloud/vector3"
)


func GenerateRandomClouds(nbClouds int, step, voxelSize, sharpness, cloudCoverVal, densityFactor float64,
                        min, max []int) []voxel_grid.VoxelGrid {
    minX := min[0]
    maxX := max[0]
    minY := min[1]
    maxY := max[1]
    minZ := min[2]
    maxZ := max[2]
    voxelGrids := make([]voxel_grid.VoxelGrid, nbClouds)

    for i := 0; i < nbClouds; i++ {
        rand.Seed(time.Now().UnixNano())

        shiftX := []float64 { float64(rand.Intn(maxX - minX) + minX), float64(rand.Intn(maxX - minX) + minX) }
        shiftZ := []float64 { float64(rand.Intn(maxZ - minZ) + minZ), float64(rand.Intn(maxZ - minZ) + minZ) }
        shiftY := []float64 { float64(rand.Intn(maxY - minY) + minY), float64(rand.Intn(maxY - minY) + minY) }

        sort.Float64s(shiftX)
        sort.Float64s(shiftZ)
        sort.Float64s(shiftY)

        perlinNoise := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.5, 3, int64(i))
        shift := vector3.InitVector3(shiftX[0], shiftY[0], shiftZ[0])
        oppositeCorner := vector3.InitVector3(shiftX[1], shiftY[1], shiftZ[1])

        voxelGrids[i] = voxel_grid.InitVoxelGrid(voxelSize, shift, oppositeCorner, step,
            perlinNoise, sharpness, cloudCoverVal, densityFactor)
        fmt.Println("shift")
        fmt.Println(shift)

        fmt.Println("opposite")
        fmt.Println(oppositeCorner)
    }

    return voxelGrids
}
