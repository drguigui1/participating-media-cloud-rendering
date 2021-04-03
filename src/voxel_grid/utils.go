package voxel_grid

import (
    "math"
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/stat/distmv"
    "golang.org/x/exp/rand"
    "volumetric-cloud/vector3"
)

func Round4(val float64) float64 {
    return math.Round(val * 10000.0) / 10000.0
}

func Round3(val float64) float64 {
    return math.Round(val * 1000.0) / 1000.0
}

func GaussianPdf(mu []float64, mat_cov []float64, seed uint64, v vector3.Vector3) (float64, bool){
    sigma := mat.NewSymDense(3, mat_cov)
    src := rand.NewSource(seed)
    normal, bool := distmv.NewNormal(mu, sigma, src)
    if !bool {
        return 0, false
    }
    return normal.Prob([]float64 {v.X, v.Y, v.Z}), true
}