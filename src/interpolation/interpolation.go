package interpolation

import (
    "math"
)

// linear 1D interpolation
func Lerp(lo, hi, t float64) float64 {
    return lo * (1 - t) + hi * t
}

func CosineInterpolate(lo, hi, t float64) float64 {
    ft := t * 3.14159265358979323846
    f := (1 - math.Cos(ft))*0.5

    return  lo *(1 - f) + hi * f;
}
