package gaussian_tower

import (
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/stat/distmv"
    "golang.org/x/exp/rand"
    r_math "math/rand"

)

type GaussianTower struct {
    mu []float64
    sigma *mat.SymDense
    seed rand.Source
}

func InitRandomGaussianTower(seed uint64, rangeX, rangeY, rangeZ []float64 ) GaussianTower{
    muX := r_math.Float64()
    muY := r_math.Float64()
    muZ := r_math.Float64()

    mu := []float64 {
                muX * (rangeX[1] - rangeX[0] + rangeX[0]),
                muY * (rangeY[1] - rangeY[0] + rangeY[0]),
                muZ * (rangeZ[1] - rangeZ[0] + rangeZ[0]),
    }
    mat_cov := make([]float64, 9)
    for i := 0; i < 9; i+=1 {
        mat_cov[i] = r_math.Float64()
    }
    mat_cov = Multiplication3(mat_cov, Transpose3(mat_cov))
    sigma := mat.NewSymDense(3, mat_cov)
    src := rand.NewSource(seed)
    return GaussianTower{mu, sigma, src}
}


func InitGaussianTower(mu []float64, mat_cov []float64, seed uint64) GaussianTower{
    sigma := mat.NewSymDense(3, mat_cov)
    src := rand.NewSource(seed)
    return GaussianTower{mu, sigma, src}
}

func (gt GaussianTower) GaussianPdf(p []float64) (float64, bool){
    normal, bool := distmv.NewNormal(gt.mu, gt.sigma, gt.seed)
    if !bool {
        return 0, false
    }
    return normal.Prob(p), true
}
