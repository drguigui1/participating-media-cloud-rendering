package height_distribution

import (
    "math"
)

func HeightDistribution(height, a, b float64) float64 {
    return (1 - math.Exp(-a * height)) * math.Exp(-b * height)
}
