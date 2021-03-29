package voxel_grid

import (
    "math"
)

func Round4(val float64) float64 {
    return math.Round(val * 10000.0) / 10000.0
}

func Round3(val float64) float64 {
    return math.Round(val * 1000.0) / 1000.0
}
