package interpolation

// linear 1D interpolation
func Lerp(lo, hi, t float64) float64 {
    return lo * (1 - t) + hi * t
}
