package atmosphere

import (
    "math"
)

func RayleighPhase(mu float64) float64 {
    return (3.0 / (16 * math.Pi)) (1 + mu * mu)
}

// g: control the anisotropy of the medium
func MiePhase(mu float64, g float64) float64 {
    var num float64 = (3 / (8 * math.Pi)) * (1 - g * g) * (1 + mu * mu) 
    var den float64 = (2 + g * g) * math.Pow((1 + g * g - 2 * g * mu), 1.5)
    return num / den
}
