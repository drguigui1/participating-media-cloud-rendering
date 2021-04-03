package height_distribution

import (
    "math"
)

func HeightDistribution(height, a, b float64) float64 {
    return (1 - math.Exp(-a * height)) * math.Exp(-b * height)
}

func GaussianTower(amp, x, y, z float64, sigma []float64, mean []float64) float64{
    return amp * (math.Exp(-(math.Pow(x - mean[0], 2.0)) / (2.0 * sigma[0]) - math.Pow(y - mean[1], 2.0) / (2.0 * sigma[1])));
}
